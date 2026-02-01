# Goose Integration

**Support Level: ✅ FULL**

This integration provides full shared skills support for Goose through:
- MCP server integration
- Native toolkit definitions
- Profile configurations

## Installation

### 1. Copy Profile Configuration

Add the profile configuration to your Goose config:

```bash
# Copy to Goose config directory
cp profiles.yaml ~/.config/goose/profiles.yaml

# Or merge with existing profiles
```

### 2. Copy Toolkit Definition

Copy the toolkit for native skill access:

```bash
mkdir -p ~/.config/goose/toolkits
cp toolkit.yaml ~/.config/goose/toolkits/shared-skills.yaml
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

### Native Toolkit Commands

When using the native toolkit:

- `shared_deploy` - Deploy skill
- `shared_test` - Test skill
- `shared_lint` - Lint skill
- `shared_format` - Format skill

## Usage Examples

### Using MCP Tools

```
User: Deploy to staging
Goose: I'll deploy to staging using the MCP deploy tool...
```

### Using Native Toolkit

```
User: Run the shared lint skill
Goose: Using shared_lint to check code quality...
```

### Using Profiles

Switch to the shared-skills profile for full access:

```bash
goose session --profile shared-skills
```

## Configuration

### profiles.yaml

Defines profiles with skill toolkits enabled.

### toolkit.yaml

Defines native skill commands that wrap the shared scripts.

## Customization

Edit toolkit.yaml to add project-specific skills or modify parameters.
