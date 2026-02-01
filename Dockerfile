# Minimal Docker image with agent CLIs for sandboxes
# Base: Debian Bookworm Slim (~25MB)
# Agents: OpenCode, Claude Code, Goose, Codex, Cursor CLI

FROM debian:bookworm-slim

LABEL org.opencontainers.image.source="https://github.com/devdone-labs/dd-agents-registry"
LABEL org.opencontainers.image.description="Minimal agent CLIs image for AI coding assistants"
LABEL org.opencontainers.image.licenses="MIT"

# Prevent interactive prompts during package installation
ENV DEBIAN_FRONTEND=noninteractive

# Base dependencies (added bzip2 for Goose binary extraction)
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    ca-certificates \
    git \
    unzip \
    bzip2 \
    && rm -rf /var/lib/apt/lists/*

# Install Node.js LTS (required for Claude Code, Codex)
RUN curl -fsSL https://deb.nodesource.com/setup_lts.x | bash - \
    && apt-get install -y --no-install-recommends nodejs \
    && rm -rf /var/lib/apt/lists/* \
    && npm cache clean --force

# =============================================================================
# Agent CLIs Installation
# =============================================================================

# OpenCode - Go-based AI coding assistant
# https://opencode.ai
# Downloads pre-built binary for the current architecture
ARG TARGETARCH
RUN set -eux; \
    case "${TARGETARCH}" in \
        amd64) OPENCODE_ARCH="amd64" ;; \
        arm64) OPENCODE_ARCH="arm64" ;; \
        *) echo "Unsupported architecture: ${TARGETARCH}" && exit 1 ;; \
    esac; \
    curl -fsSL "https://github.com/opencode-ai/opencode/releases/latest/download/opencode-linux-${OPENCODE_ARCH}" -o /usr/local/bin/opencode \
    && chmod +x /usr/local/bin/opencode \
    || echo "OpenCode download failed, trying alternative..."

# Claude Code - Anthropic's AI coding assistant
# https://docs.anthropic.com/claude-code
RUN npm install -g @anthropic-ai/claude-code \
    && npm cache clean --force

# Goose - AI developer agent (using official install script)
# https://github.com/block/goose
# CONFIGURE=false skips interactive prompts
RUN curl -fsSL https://github.com/block/goose/releases/download/stable/download_cli.sh | CONFIGURE=false bash \
    && mv /root/.local/bin/goose /usr/local/bin/goose 2>/dev/null || true \
    && chmod +x /usr/local/bin/goose

# Codex - OpenAI's coding assistant
# https://github.com/openai/codex
RUN npm install -g @openai/codex \
    && npm cache clean --force

# Cursor CLI - Cursor's AI agent
# https://cursor.com/cli
# Note: Cursor CLI requires interactive setup, may not work in headless Docker
RUN curl https://cursor.com/install -fsSL | bash \
    && if [ -f /root/.local/bin/agent ]; then \
         mv /root/.local/bin/agent /usr/local/bin/cursor-agent \
         && chmod +x /usr/local/bin/cursor-agent; \
       fi \
    || echo "Cursor CLI installation skipped (may require interactive setup)"

# =============================================================================
# Cleanup and finalization
# =============================================================================

# Create non-root user for running agents
RUN useradd -m -s /bin/bash agentuser \
    && mkdir -p /workspace \
    && chown -R agentuser:agentuser /workspace

# Set working directory
WORKDIR /workspace

# Verify installations (will FAIL build if any agent is missing)
RUN echo "=== Verifying Installed Agents ===" \
    && opencode --version \
    && echo "OpenCode: OK" \
    && claude --version \
    && echo "Claude Code: OK" \
    && goose --version \
    && echo "Goose: OK" \
    && codex --version \
    && echo "Codex: OK" \
    && cursor-agent --version \
    && echo "Cursor CLI: OK" \
    && echo "=== All agents verified successfully ==="

# Default to non-root user (can be overridden)
USER agentuser

# Default command - show available agents
CMD ["sh", "-c", "echo 'Available agents: opencode, claude, goose, codex, cursor-agent'"]
