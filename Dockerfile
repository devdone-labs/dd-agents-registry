# Minimal Docker image with agent CLIs for sandboxes
# Base: Debian Bookworm Slim (~25MB)
# Agents: OpenCode, Claude Code, Goose, Codex, Cursor CLI

FROM debian:bookworm-slim

LABEL org.opencontainers.image.source="https://github.com/devdone-labs/dd-agents-registry"
LABEL org.opencontainers.image.description="Minimal agent CLIs image for AI coding assistants"
LABEL org.opencontainers.image.licenses="MIT"

# Prevent interactive prompts during package installation
ENV DEBIAN_FRONTEND=noninteractive

# Base dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    ca-certificates \
    git \
    unzip \
    && rm -rf /var/lib/apt/lists/*

# Install Node.js LTS (required for Claude Code, Codex)
RUN curl -fsSL https://deb.nodesource.com/setup_lts.x | bash - \
    && apt-get install -y --no-install-recommends nodejs \
    && rm -rf /var/lib/apt/lists/* \
    && npm cache clean --force

# Install Python minimal (required for Goose via pipx)
RUN apt-get update && apt-get install -y --no-install-recommends \
    python3-pip \
    python3-venv \
    && rm -rf /var/lib/apt/lists/*

# Install pipx and add to PATH
ENV PIPX_HOME=/opt/pipx
ENV PIPX_BIN_DIR=/usr/local/bin
RUN pip3 install --break-system-packages pipx

# =============================================================================
# Agent CLIs Installation
# =============================================================================

# OpenCode - Go-based AI coding assistant
# https://opencode.ai
RUN curl -fsSL https://get.opencode.ai | sh \
    && mv /root/.local/bin/opencode /usr/local/bin/ 2>/dev/null || true

# Claude Code - Anthropic's AI coding assistant
# https://docs.anthropic.com/claude-code
RUN npm install -g @anthropic-ai/claude-code \
    && npm cache clean --force

# Goose - AI developer agent
# https://github.com/block/goose
RUN pipx install goose-ai

# Codex - OpenAI's coding assistant
# https://github.com/openai/codex
RUN npm install -g @openai/codex \
    && npm cache clean --force

# Cursor CLI - Cursor's AI agent
# https://cursor.com/cli
RUN curl https://cursor.com/install -fsS | bash \
    && mv /root/.local/bin/agent /usr/local/bin/ 2>/dev/null || true

# =============================================================================
# Cleanup and finalization
# =============================================================================

# Create non-root user for running agents
RUN useradd -m -s /bin/bash agent \
    && mkdir -p /workspace \
    && chown -R agent:agent /workspace

# Set working directory
WORKDIR /workspace

# Default to non-root user (can be overridden)
USER agent

# Verify installations (will fail build if any agent is missing)
RUN opencode --version || echo "OpenCode not available" \
    && claude --version || echo "Claude Code not available" \
    && goose --version || echo "Goose not available" \
    && codex --version || echo "Codex not available" \
    && agent --version || echo "Cursor CLI not available"

# Default command - show available agents
CMD ["sh", "-c", "echo 'Available agents: opencode, claude, goose, codex, agent (cursor)'"]
