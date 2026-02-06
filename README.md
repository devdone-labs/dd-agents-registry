# DD Agents Registry

Minimal Docker images containing AI coding agent CLIs for use in sandboxed development environments.

## Available Agents

| Agent | CLI Command | Description |
|-------|-------------|-------------|
| [OpenCode](https://opencode.ai) | `opencode` | Go-based AI coding assistant |
| [Claude Code](https://docs.anthropic.com/claude-code) | `claude` | Anthropic's AI coding assistant |
| [Goose](https://github.com/block/goose) | `goose` | AI developer agent by Block |
| [Codex](https://github.com/openai/codex) | `codex` | OpenAI's coding assistant |
| [Cursor CLI](https://cursor.com/cli) | `cursor-agent` | Cursor's AI agent |

## Quick Start

### Linux

#### Pull the Image

```bash
docker pull ghcr.io/devdone-labs/dd-agents:latest
```

#### Run an Agent

```bash
# Run OpenCode
docker run -it --rm -v $(pwd):/workspace ghcr.io/devdone-labs/dd-agents:latest opencode

# Run Claude Code
docker run -it --rm -v $(pwd):/workspace ghcr.io/devdone-labs/dd-agents:latest claude

# Run Goose
docker run -it --rm -v $(pwd):/workspace ghcr.io/devdone-labs/dd-agents:latest goose

# Run Codex
docker run -it --rm -v $(pwd):/workspace ghcr.io/devdone-labs/dd-agents:latest codex

# Run Cursor CLI
docker run -it --rm -v $(pwd):/workspace ghcr.io/devdone-labs/dd-agents:latest cursor-agent
```

### Windows

> **Note:** Windows containers require Windows 10/11 Pro, Enterprise, or Windows Server with the Containers feature enabled.

#### Pull the Image

```powershell
docker pull ghcr.io/devdone-labs/dd-agents-windows:latest
```

#### Run an Agent

```powershell
# Run OpenCode
docker run -it --rm -v ${PWD}:C:\workspace ghcr.io/devdone-labs/dd-agents-windows:latest opencode

# Run Claude Code
docker run -it --rm -v ${PWD}:C:\workspace ghcr.io/devdone-labs/dd-agents-windows:latest claude

# Run Goose
docker run -it --rm -v ${PWD}:C:\workspace ghcr.io/devdone-labs/dd-agents-windows:latest goose

# Run Codex
docker run -it --rm -v ${PWD}:C:\workspace ghcr.io/devdone-labs/dd-agents-windows:latest codex

# Run Cursor CLI
docker run -it --rm -v ${PWD}:C:\workspace ghcr.io/devdone-labs/dd-agents-windows:latest cursor-agent
```

### With API Keys

Most agents require API keys. Pass them as environment variables:

```bash
# Linux/macOS
docker run -it --rm \
  -e ANTHROPIC_API_KEY=$ANTHROPIC_API_KEY \
  -e OPENAI_API_KEY=$OPENAI_API_KEY \
  -v $(pwd):/workspace \
  ghcr.io/devdone-labs/dd-agents:latest claude
```

```powershell
# Windows
docker run -it --rm `
  -e ANTHROPIC_API_KEY=$env:ANTHROPIC_API_KEY `
  -e OPENAI_API_KEY=$env:OPENAI_API_KEY `
  -v ${PWD}:C:\workspace `
  ghcr.io/devdone-labs/dd-agents-windows:latest claude
```

## Image Details

### Linux Image

- **Base:** Debian Bookworm Slim (~25MB base)
- **Platforms:** linux/amd64, linux/arm64
- **Registry:** `ghcr.io/devdone-labs/dd-agents`

| Tag | Description |
|-----|-------------|
| `latest` | Latest stable build from main branch |
| `v1.0.0` | Specific version release |
| `<sha>` | Specific commit SHA |

### Windows Image

- **Base:** Windows Nano Server LTSC 2022 (~100MB base)
- **Platforms:** windows/amd64
- **Registry:** `ghcr.io/devdone-labs/dd-agents-windows` (separate package)

| Tag | Description |
|-----|-------------|
| `latest` | Latest stable Windows build from main branch |
| `ltsc2022` | Windows Server 2022 LTSC build |
| `v1.0.0` | Specific version release |
| `<sha>` | Specific commit SHA |

## Runtime Installation

Language runtimes (Go, Python, Rust, etc.) are **not** included in the base images to keep them minimal. Agents can install required runtimes on-demand within sandboxes using skills.

## Development

### Build Locally

```bash
# Linux
docker build -t dd-agents:local .

# Windows (requires Windows host with Docker in Windows container mode)
docker build -f Dockerfile.windows -t dd-agents:windows-local .
```

### Test Agents

```bash
# Linux
docker run --rm dd-agents:local opencode --version
docker run --rm dd-agents:local claude --version
docker run --rm dd-agents:local goose --version
docker run --rm dd-agents:local codex --version
docker run --rm dd-agents:local cursor-agent --version
```

```powershell
# Windows
docker run --rm dd-agents:windows-local opencode --version
docker run --rm dd-agents:windows-local claude --version
docker run --rm dd-agents:windows-local goose --version
docker run --rm dd-agents:windows-local codex --version
docker run --rm dd-agents:windows-local cursor-agent --version
```

## CI/CD

This repository uses GitHub Actions to automatically build and push Docker images to GitHub Container Registry on:

- Push to `main` branch
- Version tags (`v*`)
- Manual workflow dispatch

Both Linux and Windows images are built in parallel:
- **Linux:** Built on `ubuntu-latest` with multi-platform support (amd64, arm64)
- **Windows:** Built on `windows-2022` for Windows Server 2022 LTSC

## License

MIT
