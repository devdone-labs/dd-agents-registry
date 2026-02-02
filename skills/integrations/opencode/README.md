# OpenCode Integration

**Support Level: ⚠️ LIMITED**

This integration provides basic shared skills support for OpenCode through:
- Basic MCP server configuration
- Config-based instructions
- No native plugin system

## Limitations

OpenCode has limited skill support:
- ⚠️ Basic MCP support (may be incomplete)
- ✅ Configuration file instructions
- ❌ No native executable skills
- ❌ No toolkit/plugin system
- ❌ No slash commands

## Installation

### 1. Copy Configuration

Copy the config file to your OpenCode directory:

```bash
cp config.json ~/.opencode/config.json
```

### 2. Copy Instructions

Copy the instructions file to your project:

```bash
cp INSTRUCTIONS.md /path/to/your/project/INSTRUCTIONS.md
```

### 3. Install MCP Server Dependencies

```bash
cd ../../shared/mcp-server
npm install
```

## Features

### MCP Tools (Basic Support)

| Tool | Description | Support |
|------|-------------|---------|
| `deploy` | Deploy to environments | ⚠️ Basic |
| `test` | Run test suite | ⚠️ Basic |
| `lint` | Run linters | ⚠️ Basic |
| `format` | Format code | ⚠️ Basic |

### Manual Skill Execution

For best results, use the shared scripts directly:

```bash
# Deploy
./skills/shared/scripts/deploy-execute.sh --env=staging

# Test
./skills/shared/scripts/test-run.sh --coverage

# Lint
./skills/shared/scripts/lint-run.sh --fix

# Format
./skills/shared/scripts/format-run.sh
```

## Usage

### With MCP (if supported)

```
User: Deploy to staging
OpenCode: Using deploy tool...
```

### Manual Execution

```
User: Deploy to staging
OpenCode: I'll run the deployment script:
./skills/shared/scripts/deploy-execute.sh --env=staging
```

## Configuration

### config.json

Contains MCP server configuration and project instructions.

### INSTRUCTIONS.md

Contains skill documentation for the agent to follow.
