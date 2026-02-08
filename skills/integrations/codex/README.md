# Codex Integration

**Support Level: ⚠️ LIMITED**

This integration provides shared skills support for OpenAI Codex through:
- MCP server integration
- AGENTS.md instruction file
- No native plugin system (instruction-based only)

## Limitations

Codex has limited skill support compared to Claude Code and Goose:
- ✅ MCP tools work via server integration
- ✅ Instructions in AGENTS.md are understood
- ❌ No native executable skills
- ❌ No slash commands
- ❌ No toolkit/plugin system

## Installation

### 1. Copy Instructions File

Copy the AGENTS.md file to your project:

```bash
cp AGENTS.md /path/to/your/project/AGENTS.md
```

### 2. Configure MCP (if supported)

Add MCP server to your Codex configuration:

```bash
cp config.yaml ~/.codex/config.yaml
```

### 3. Install MCP Server Dependencies

```bash
cd ../../shared/mcp-server
npm install
```

## Features

### MCP Tools Available

| Tool | Description |
|------|-------------|
| `deploy` | Deploy to dev/staging/prod environments |
| `test` | Run test suite with coverage options |
| `lint` | Run linters with auto-fix |
| `format` | Format code with project formatters |

### Instruction-Based Skills

AGENTS.md provides documentation that Codex can follow for:
- Deployment procedures
- Testing guidelines
- Linting rules
- Formatting standards

## Usage Examples

### Using MCP Tools

```
User: Deploy to staging
Codex: I'll use the deploy MCP tool to deploy to staging...
```

### Instruction-Based Execution

```
User: Run the deployment process
Codex: Following the AGENTS.md instructions, I'll:
1. Validate configuration
2. Build artifacts
3. Execute deployment
```

## Configuration

### config.yaml

```yaml
mcp:
  servers:
    shared-skills:
      command: node
      args:
        - ./skills/shared/mcp-server/index.js
```

### AGENTS.md

Contains skill documentation and step-by-step instructions for manual execution.
