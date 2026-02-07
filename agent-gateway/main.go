// agent-gateway is a lightweight HTTP server that runs as PID 1 inside a
// libkrun VM. It receives messages from the host via HTTP, spawns the
// appropriate agent CLI, and streams output back as Server-Sent Events (SSE).
//
// This enables persistent VMs with multi-turn conversations: the VM boots
// once, and each chat message is an HTTP request rather than a new VM.
package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// ---------------------------------------------------------------------------
// Configuration
// ---------------------------------------------------------------------------

const (
	listenAddr     = ":8080"
	defaultTimeout = 300 * time.Second
	defaultWorkDir = "/workspace"
)

// ---------------------------------------------------------------------------
// Types
// ---------------------------------------------------------------------------

// MessageRequest is the JSON body for POST /api/v1/message.
type MessageRequest struct {
	Message string            `json:"message"`
	Agent   string            `json:"agent"`
	Model   string            `json:"model"`
	Env     map[string]string `json:"env"`
	Timeout int               `json:"timeout"` // seconds, 0 = default
}

// ExecRequest is the JSON body for POST /api/v1/exec.
type ExecRequest struct {
	Command string            `json:"command"`
	Args    []string          `json:"args"`
	Env     map[string]string `json:"env"`
	Timeout int               `json:"timeout"`
}

// SSEEvent is a single SSE data payload.
type SSEEvent struct {
	Type string `json:"type"` // "stdout", "stderr", "exit", "error"
	Data string `json:"data,omitempty"`
	Code *int   `json:"code,omitempty"`
}

// agentSession tracks per-agent state for conversation continuity.
type agentSession struct {
	messageCount int
	sessionID    string // captured from agent output (goose, opencode)
}

// ---------------------------------------------------------------------------
// Global state
// ---------------------------------------------------------------------------

var (
	startTime = time.Now()
	busy      sync.Mutex // only one agent command at a time
	sessions  = struct {
		sync.Mutex
		m map[string]*agentSession
	}{m: make(map[string]*agentSession)}
)

func getSession(agent string) *agentSession {
	sessions.Lock()
	defer sessions.Unlock()
	s, ok := sessions.m[agent]
	if !ok {
		s = &agentSession{}
		sessions.m[agent] = s
	}
	return s
}

// ---------------------------------------------------------------------------
// Network initialisation (replaces nanosb-net-init shell script)
// ---------------------------------------------------------------------------

func initNetwork() {
	// Configure eth0 with the static IP that gvproxy expects.
	// Try `ip` first (iproute2), fall back to `ifconfig`/`route`.
	if path, err := exec.LookPath("ip"); err == nil {
		run(path, "link", "set", "eth0", "up")
		run(path, "addr", "add", "192.168.127.2/24", "dev", "eth0")
		run(path, "route", "add", "default", "via", "192.168.127.1", "dev", "eth0")
	} else if path, err := exec.LookPath("ifconfig"); err == nil {
		run(path, "eth0", "192.168.127.2", "netmask", "255.255.255.0", "up")
		if routePath, err := exec.LookPath("route"); err == nil {
			run(routePath, "add", "default", "gw", "192.168.127.1")
		}
	} else {
		log.Println("[agent-gateway] WARNING: no networking tools found (ip/ifconfig)")
	}

	// Configure DNS — gvproxy's built-in DNS is at the gateway IP.
	_ = os.MkdirAll("/etc", 0755)
	_ = os.WriteFile("/etc/resolv.conf", []byte("nameserver 192.168.127.1\n"), 0644)
}

func run(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Printf("[agent-gateway] %s %v: %v", name, args, err)
	}
}

// ---------------------------------------------------------------------------
// Agent command builder
// ---------------------------------------------------------------------------

func buildAgentCommand(req *MessageRequest, sess *agentSession) (string, []string) {
	switch req.Agent {
	case "claude", "claude-code":
		args := []string{
			"--print", req.Message,
			"--verbose",
			"--output-format", "stream-json",
			"--include-partial-messages",
		}
		if req.Model != "" {
			args = append(args, "--model", req.Model)
		}
		// Add --continue for 2nd+ message to resume conversation context.
		// Claude Code stores sessions in ~/.claude/ which persists in the VM.
		if sess.messageCount > 0 {
			args = append(args, "--continue")
		}
		return "claude", args

	case "opencode":
		args := []string{"chat", "--message", req.Message}
		if sess.sessionID != "" {
			args = append(args, "--session", sess.sessionID)
		}
		return "opencode", args

	case "goose":
		if sess.messageCount > 0 && sess.sessionID != "" {
			return "goose", []string{"session", "resume", sess.sessionID, "--message", req.Message}
		}
		return "goose", []string{"run", "--text", req.Message}

	case "codex":
		args := []string{}
		if req.Model != "" {
			args = append(args, "--model", req.Model)
		}
		args = append(args, req.Message)
		return "codex", args

	case "cursor", "cursor-agent":
		return "cursor-agent", []string{"--message", req.Message}

	default:
		// Custom/unknown agent: treat agent name as the binary
		return req.Agent, []string{req.Message}
	}
}

// ---------------------------------------------------------------------------
// SSE streaming helpers
// ---------------------------------------------------------------------------

func sseWrite(w http.ResponseWriter, evt SSEEvent) {
	data, _ := json.Marshal(evt)
	fmt.Fprintf(w, "data: %s\n\n", data)
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}

// streamCommand spawns a command, streams stdout/stderr as SSE, and returns
// the exit code. It blocks until the command finishes or the context is done.
func streamCommand(ctx context.Context, w http.ResponseWriter, bin string, args []string, env map[string]string) int {
	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.Dir = defaultWorkDir

	// Build environment: inherit base env, overlay request-specific vars.
	cmdEnv := os.Environ()
	for k, v := range env {
		cmdEnv = append(cmdEnv, fmt.Sprintf("%s=%s", k, v))
	}
	// Ensure basic vars are set.
	cmdEnv = ensureEnv(cmdEnv, "HOME", "/root")
	cmdEnv = ensureEnv(cmdEnv, "PATH", "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin")
	cmdEnv = ensureEnv(cmdEnv, "TERM", "dumb")
	cmd.Env = cmdEnv

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		sseWrite(w, SSEEvent{Type: "error", Data: fmt.Sprintf("stdout pipe: %v", err)})
		return -1
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		sseWrite(w, SSEEvent{Type: "error", Data: fmt.Sprintf("stderr pipe: %v", err)})
		return -1
	}

	if err := cmd.Start(); err != nil {
		sseWrite(w, SSEEvent{Type: "error", Data: fmt.Sprintf("start: %v", err)})
		return -1
	}

	// Stream stdout and stderr concurrently.
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stdout)
		scanner.Buffer(make([]byte, 256*1024), 1024*1024) // 1MB line buffer
		for scanner.Scan() {
			sseWrite(w, SSEEvent{Type: "stdout", Data: scanner.Text()})
		}
	}()

	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stderr)
		scanner.Buffer(make([]byte, 256*1024), 1024*1024)
		for scanner.Scan() {
			sseWrite(w, SSEEvent{Type: "stderr", Data: scanner.Text()})
		}
	}()

	wg.Wait()

	exitCode := 0
	if err := cmd.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			exitCode = -1
		}
	}
	return exitCode
}

func ensureEnv(env []string, key, fallback string) []string {
	prefix := key + "="
	for _, e := range env {
		if len(e) > len(prefix) && e[:len(prefix)] == prefix {
			return env
		}
	}
	return append(env, prefix+fallback)
}

// ---------------------------------------------------------------------------
// HTTP Handlers
// ---------------------------------------------------------------------------

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	uptimeMs := time.Since(startTime).Milliseconds()
	fmt.Fprintf(w, `{"status":"ok","uptime_ms":%d}`, uptimeMs)
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Only one agent command at a time.
	if !busy.TryLock() {
		http.Error(w, `{"error":"agent is busy with another request"}`, http.StatusConflict)
		return
	}
	defer busy.Unlock()

	var req MessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"invalid json: %v"}`, err), http.StatusBadRequest)
		return
	}

	sess := getSession(req.Agent)
	bin, args := buildAgentCommand(&req, sess)

	timeout := defaultTimeout
	if req.Timeout > 0 {
		timeout = time.Duration(req.Timeout) * time.Second
	}
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	// SSE headers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")
	w.WriteHeader(http.StatusOK)

	log.Printf("[agent-gateway] message #%d to %s: %s %v",
		sess.messageCount+1, req.Agent, bin, args)

	exitCode := streamCommand(ctx, w, bin, args, req.Env)

	code := exitCode
	sseWrite(w, SSEEvent{Type: "exit", Code: &code})

	sess.messageCount++
}

func execHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !busy.TryLock() {
		http.Error(w, `{"error":"agent is busy with another request"}`, http.StatusConflict)
		return
	}
	defer busy.Unlock()

	var req ExecRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"invalid json: %v"}`, err), http.StatusBadRequest)
		return
	}

	timeout := defaultTimeout
	if req.Timeout > 0 {
		timeout = time.Duration(req.Timeout) * time.Second
	}
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")
	w.WriteHeader(http.StatusOK)

	log.Printf("[agent-gateway] exec: %s %v", req.Command, req.Args)

	exitCode := streamCommand(ctx, w, req.Command, req.Args, req.Env)

	code := exitCode
	sseWrite(w, SSEEvent{Type: "exit", Code: &code})
}

func stopHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status":"stopping"}`)
	log.Println("[agent-gateway] stop requested, shutting down")
	go func() {
		time.Sleep(100 * time.Millisecond)
		os.Exit(0)
	}()
}

// ---------------------------------------------------------------------------
// Main
// ---------------------------------------------------------------------------

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("[agent-gateway] starting...")

	// As PID 1 we must reap orphan children to avoid zombies.
	// We do NOT ignore SIGCHLD because Go's exec.Cmd.Wait() uses waitpid()
	// and ignoring SIGCHLD causes the kernel to auto-reap, making Wait() fail
	// with "no child processes".
	// Instead, start a goroutine that reaps any orphan processes.
	go func() {
		for {
			var ws syscall.WaitStatus
			pid, err := syscall.Wait4(-1, &ws, syscall.WNOHANG, nil)
			if pid <= 0 || err != nil {
				time.Sleep(1 * time.Second)
			}
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	// Configure networking before starting the HTTP server.
	initNetwork()
	log.Println("[agent-gateway] network configured")

	// Ensure working directory exists.
	_ = os.MkdirAll(defaultWorkDir, 0755)

	// Register routes.
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/api/v1/health", healthHandler)
	mux.HandleFunc("/api/v1/message", messageHandler)
	mux.HandleFunc("/api/v1/exec", execHandler)
	mux.HandleFunc("/api/v1/stop", stopHandler)

	server := &http.Server{
		Addr:         listenAddr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 0, // SSE streams can be long-lived
	}

	// Start HTTP server in background.
	go func() {
		log.Printf("[agent-gateway] listening on %s", listenAddr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("[agent-gateway] server error: %v", err)
		}
	}()

	// Wait for shutdown signal.
	sig := <-sigCh
	log.Printf("[agent-gateway] received %v, shutting down", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = server.Shutdown(ctx)
}
