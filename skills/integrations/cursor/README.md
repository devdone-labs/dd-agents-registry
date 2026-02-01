# Cursor CLI Integration

**Support Level: ❌ MINIMAL**

This integration provides minimal shared skills support for Cursor CLI through:
- `.cursorrules` documentation only
- No MCP support in CLI mode
- No executable skills

## Limitations

Cursor CLI has minimal standalone skill support:
- ❌ No MCP support (requires IDE)
- ❌ No native executable skills
- ❌ No toolkit/plugin system
- ✅ Rules files for documentation only

**Note:** Full Cursor skill features require the Cursor IDE, not the CLI.

## What This Integration Provides

Since Cursor CLI cannot execute skills directly, this integration provides:
1. Documentation in `.cursorrules` format
2. Instructions for manual skill execution
3. Guidance for the agent to follow

## Installation

### Copy Rules File

Copy the rules file to your project:

```bash
cp .cursorrules /path/to/your/project/.cursorrules
```

## Features

### Documentation Only

The `.cursorrules` file documents available skills so Cursor CLI can:
- Understand what skills exist
- Guide users to manual execution
- Provide context about project workflows

### Manual Skill Execution

Users must run skills manually:

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

## For Full Skill Support

To get full skill support with Cursor:

1. Use **Cursor IDE** instead of CLI
2. Configure MCP servers in IDE settings
3. Use the IDE's extension system

### IDE MCP Configuration

In Cursor IDE settings, add:

```json
{
  "mcpServers": {
    "shared-skills": {
      "command": "node",
      "args": ["./skills/shared/mcp-server/index.js"]
    }
  }
}
```

## Comparison

| Feature | Cursor CLI | Cursor IDE |
|---------|------------|------------|
| Rules files | ✅ | ✅ |
| MCP servers | ❌ | ✅ |
| Executable skills | ❌ | ✅ |
| Extensions | ❌ | ✅ |
