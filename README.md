# DD Agents Registry

Minimal Docker image containing AI coding agent CLIs for use in sandboxed development environments.

## Available Agents

| Agent | CLI Command | Description |
|-------|-------------|-------------|
| [OpenCode](https://opencode.ai) | `opencode` | Go-based AI coding assistant |
| [Claude Code](https://docs.anthropic.com/claude-code) | `claude` | Anthropic's AI coding assistant |
| [Goose](https://github.com/block/goose) | `goose` | AI developer agent by Block |
| [Codex](https://github.com/openai/codex) | `codex` | OpenAI's coding assistant |
| [Cursor CLI](https://cursor.com/cli) | `cursor-agent` | Cursor's AI agent |

## Quick Start

### Pull the Image

```bash
docker pull ghcr.io/devdone-labs/dd-agents:latest
```

### Run an Agent

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

### With API Keys

Most agents require API keys. Pass them as environment variables:

```bash
docker run -it --rm \
  -e ANTHROPIC_API_KEY=$ANTHROPIC_API_KEY \
  -e OPENAI_API_KEY=$OPENAI_API_KEY \
  -v $(pwd):/workspace \
  ghcr.io/devdone-labs/dd-agents:latest claude
```

## Image Details

- **Base:** Debian Bookworm Slim (~25MB base)
- **Platforms:** linux/amd64, linux/arm64
- **Registry:** `ghcr.io/devdone-labs/dd-agents`

### Tags

| Tag | Description |
|-----|-------------|
| `latest` | Latest stable build from main branch |
| `v1.0.0` | Specific version release |
| `<sha>` | Specific commit SHA |

## Runtime Installation

Language runtimes (Go, Python, Rust, etc.) are **not** included in the base image to keep it minimal. Agents can install required runtimes on-demand within sandboxes using skills.

## Development

### Build Locally

```bash
docker build -t dd-agents:local .
```

### Test Agents

```bash
docker run --rm dd-agents:local opencode --version
docker run --rm dd-agents:local claude --version
docker run --rm dd-agents:local goose --version
docker run --rm dd-agents:local codex --version
docker run --rm dd-agents:local cursor-agent --version
```

## CI/CD

This repository uses GitHub Actions to automatically build and push Docker images to GitHub Container Registry on:

- Push to `main` branch
- Version tags (`v*`)
- Manual workflow dispatch

## License

MIT
