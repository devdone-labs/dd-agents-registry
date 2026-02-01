# Skills Implementation Plan

## Overview

This document outlines the architecture for a unified skills system with per-agent integrations.

## Architecture

```
skills/
├── core/                    # Universal skill definitions
│   ├── deploy/
│   ├── test/
│   ├── lint/
│   └── ...
├── integrations/            # Per-agent adapters
│   ├── claude-code/         # Full integration
│   ├── goose/               # Full integration
│   ├── codex/               # Limited integration
│   ├── opencode/            # Limited integration
│   └── cursor/              # Minimal integration
└── shared/                  # Shared utilities
    ├── mcp-server/          # Universal MCP server
    └── scripts/             # Common scripts
```

## Integration Capability Matrix

| Agent | MCP Server | Native Skills | Instructions | Rules Only |
|-------|------------|---------------|--------------|------------|
| Claude Code | ✅ | ✅ Commands | ✅ CLAUDE.md | - |
| Goose | ✅ | ✅ Toolkits | ✅ | - |
| Codex | ✅ | ❌ | ✅ AGENTS.md | - |
| OpenCode | ⚠️ Basic | ❌ | ✅ Config | - |
| Cursor CLI | ❌ | ❌ | ❌ | ✅ .cursorrules |

## Implementation Strategy

### Phase 1: Core Skills Definition
Define skills in a universal format that can be adapted to each agent.

### Phase 2: MCP Server (Primary Integration)
Build a shared MCP server that works with Claude Code, Goose, and Codex.

### Phase 3: Native Integrations
Create agent-specific adapters for native skill formats.

### Phase 4: Fallback Integrations
Create instruction-based fallbacks for agents with limited support.

## Skill Definition Format

Each skill in `core/` follows this structure:

```yaml
# skill.yaml
name: deploy
version: 1.0.0
description: Deploy application to target environment

inputs:
  environment:
    type: string
    enum: [dev, staging, prod]
    required: true
  dry_run:
    type: boolean
    default: false

commands:
  - name: validate
    run: ./scripts/validate.sh
  - name: build
    run: ./scripts/build.sh
  - name: deploy
    run: ./scripts/deploy.sh ${environment}

outputs:
  - deployment_url
  - version
```

## Per-Agent Integration Details

### Claude Code (Full Support)
- MCP server registration in settings.json
- CLAUDE.md with skill documentation
- Custom slash commands
- Hooks for automation

### Goose (Full Support)
- MCP server integration
- Native toolkit definitions
- Profile configurations

### Codex (Limited Support)
- MCP server integration
- AGENTS.md instruction file
- No native skill execution

### OpenCode (Limited Support)
- Basic MCP server config
- Instructions in config files
- Manual skill execution

### Cursor CLI (Minimal Support)
- .cursorrules with skill documentation
- No executable skills
- Manual guidance only
