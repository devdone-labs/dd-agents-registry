# Shared Skills for AI Coding Agents

This directory contains a unified skills system with per-agent integrations.

## Structure

```
skills/
├── core/                    # Universal skill definitions
│   ├── deploy/              # Deployment skill
│   ├── test/                # Testing skill
│   ├── lint/                # Linting skill
│   └── format/              # Code formatting skill
├── integrations/            # Per-agent adapters
│   ├── claude-code/         # Full integration (MCP + commands)
│   ├── goose/               # Full integration (MCP + toolkits)
│   ├── codex/               # Limited (MCP + instructions)
│   ├── opencode/            # Limited (config + instructions)
│   └── cursor/              # Minimal (rules only)
└── shared/                  # Shared utilities
    ├── mcp-server/          # Universal MCP server
    └── scripts/             # Common executable scripts
```

## Agent Support Levels

| Agent | Integration Level | Features |
|-------|------------------|----------|
| Claude Code | ✅ Full | MCP server, slash commands, hooks |
| Goose | ✅ Full | MCP server, native toolkits |
| Codex | ⚠️ Limited | MCP server, instruction files |
| OpenCode | ⚠️ Limited | Basic MCP, config files |
| Cursor CLI | ❌ Minimal | Rules file only |

## Quick Start

### 1. Install Shared MCP Server

```bash
cd skills/shared/mcp-server
npm install
```

### 2. Configure Your Agent

Copy the appropriate integration files to your agent's config directory:

**Claude Code:**
```bash
cp skills/integrations/claude-code/settings.json ~/.claude/
cp skills/integrations/claude-code/CLAUDE.md ~/project/
```

**Goose:**
```bash
cp skills/integrations/goose/profiles.yaml ~/.config/goose/
```

**Codex:**
```bash
cp skills/integrations/codex/AGENTS.md ~/project/
```

### 3. Use Skills

Each agent can now access the shared skills through their native interface.

## Adding New Skills

1. Create skill definition in `core/<skill-name>/skill.yaml`
2. Add implementation scripts in `shared/scripts/`
3. Register in MCP server: `shared/mcp-server/tools/`
4. Update agent integrations as needed

## Available Skills

| Skill | Description | Agents Supported |
|-------|-------------|------------------|
| deploy | Deploy to environment | Claude, Goose, Codex |
| test | Run test suite | Claude, Goose, Codex |
| lint | Run linters | Claude, Goose, Codex |
| format | Format code | Claude, Goose, Codex |
