# Claude Code Integration

**Support Level: ✅ FULL**

This integration provides full shared skills support for Claude Code through:
- MCP server for tool execution
- CLAUDE.md for documentation and commands
- Slash commands for quick access

## Installation

### 1. Copy MCP Settings

Add the MCP server configuration to your Claude Code settings:

```bash
# Global settings (all projects)
cp settings.json ~/.claude/settings.json

# Or merge with existing settings
```

### 2. Copy Project Documentation

Copy the CLAUDE.md file to your project root:

```bash
cp CLAUDE.md /path/to/your/project/CLAUDE.md
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

### Slash Commands

The CLAUDE.md file documents these quick commands:

- `/deploy <env>` - Quick deploy shortcut
- `/test` - Run all tests
- `/lint` - Check code quality
- `/format` - Format all code

## Usage Examples

### Using MCP Tools

Claude Code will automatically have access to the shared skills via MCP:

```
User: Deploy to staging
Claude: I'll use the deploy tool to deploy to staging...
[Calls deploy tool with environment="staging"]
```

### Using Slash Commands

```
/deploy prod
```

## Configuration

### settings.json

```json
{
  "mcpServers": {
    "shared-skills": {
      "command": "node",
      "args": ["/path/to/skills/shared/mcp-server/index.js"]
    }
  }
}
```

### CLAUDE.md Customization

Edit CLAUDE.md to add project-specific skills or modify existing ones.
