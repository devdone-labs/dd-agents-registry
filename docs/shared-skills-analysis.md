# Shared Skills Analysis: AI Coding Agents Comparison

This document analyzes the shared skills capabilities of AI coding agents, comparing which support skill sharing mechanisms and which do not.

## Overview

"Shared skills" in the context of AI coding agents refers to:
- **Reusable automation scripts** that can be shared across projects/users
- **Custom tools/extensions** that extend agent capabilities
- **Plugin systems** for third-party integrations
- **MCP (Model Context Protocol)** support for standardized tool interfaces
- **Command definitions** that can be distributed and reused

## Agent Comparison Matrix

| Agent | Shared Skills Support | MCP Support | Custom Tools | Skill Format | Distribution Method |
|-------|----------------------|-------------|--------------|--------------|---------------------|
| Claude Code | ✅ Full | ✅ Yes | ✅ Yes | Markdown + MCP | Local files, npm packages |
| Goose | ✅ Full | ✅ Yes | ✅ Yes | YAML toolkits | Built-in + community |
| Codex | ⚠️ Limited | ✅ Yes | ✅ Yes | Instructions + MCP | Local files |
| OpenCode | ⚠️ Limited | ✅ Yes | ⚠️ Basic | Config-based | Local only |
| Cursor CLI | ❌ Minimal | ⚠️ Partial | ❌ No | N/A | N/A |

---

## Detailed Analysis

### 1. Claude Code (Anthropic)

**Shared Skills Support: ✅ FULL**

Claude Code has robust shared skills support through multiple mechanisms:

#### Skill Mechanisms:
- **CLAUDE.md files**: Project-level instructions and skills defined in markdown
- **MCP Servers**: Full Model Context Protocol support for external tools
- **Custom Commands**: Slash commands that can be defined and shared
- **Hooks**: Lifecycle hooks for automation (pre/post commit, etc.)

#### Skill Sharing Methods:
```
~/.claude/
├── CLAUDE.md           # Global skills/instructions
├── settings.json       # MCP server configurations
└── commands/           # Custom slash commands
```

#### Key Features:
- Hierarchical skill loading (global → project → directory)
- MCP protocol for standardized tool integration
- npm-distributable MCP servers
- Built-in skill library access

#### Example Skill Definition:
```markdown
# CLAUDE.md

## Available Commands
- /deploy - Run deployment pipeline
- /test - Execute test suite with coverage

## Coding Standards
- Use TypeScript strict mode
- Follow ESLint rules
```

---

### 2. Goose (Block)

**Shared Skills Support: ✅ FULL**

Goose has the most mature and flexible skill system:

#### Skill Mechanisms:
- **Toolkits**: Modular skill packages (built-in, community, custom)
- **Extensions**: Plugin system for custom functionality
- **MCP Support**: Full Model Context Protocol integration
- **Profiles**: Shareable agent configurations

#### Built-in Toolkits:
| Toolkit | Description |
|---------|-------------|
| developer | File editing, shell commands, code analysis |
| github | GitHub operations (PRs, issues, repos) |
| jira | Jira ticket management |
| google_drive | Google Drive file access |
| memory | Persistent memory across sessions |
| computercontroller | Desktop automation |

#### Skill Sharing Methods:
```
~/.config/goose/
├── profiles.yaml       # Agent profiles with toolkit configs
├── toolkits/           # Custom toolkits directory
└── extensions/         # Third-party extensions
```

#### Example Profile with Skills:
```yaml
# profiles.yaml
default:
  provider: anthropic
  toolkits:
    - developer
    - github
    - name: custom-deploy
      path: ~/.config/goose/toolkits/deploy

custom-project:
  toolkits:
    - developer
    - jira
    - memory
```

#### Key Features:
- Hot-swappable toolkits during sessions
- Community toolkit repository
- MCP server integration
- Cross-session memory skills

---

### 3. Codex (OpenAI)

**Shared Skills Support: ⚠️ LIMITED**

Codex supports some skill mechanisms but with limitations:

#### Skill Mechanisms:
- **Instructions files**: AGENTS.md / CODEX.md for project instructions
- **MCP Support**: Model Context Protocol for external tools
- **Sandbox configs**: Environment-specific configurations

#### Skill Sharing Methods:
```
project/
├── AGENTS.md           # Project-level instructions
├── CODEX.md            # Alternative instruction file
└── .codex/
    └── config.yaml     # MCP and settings
```

#### Limitations:
- No official plugin/extension system
- Skills are instruction-based rather than executable
- Limited community skill sharing infrastructure
- MCP support is functional but less documented

#### Example Instruction-Based Skill:
```markdown
# AGENTS.md

## Project Context
This is a Python Django project.

## Skills
When asked to deploy:
1. Run `python manage.py collectstatic`
2. Run tests with `pytest`
3. Deploy using `./scripts/deploy.sh`
```

---

### 4. OpenCode

**Shared Skills Support: ⚠️ LIMITED**

OpenCode has basic configuration-based skill support:

#### Skill Mechanisms:
- **Config files**: JSON/YAML configuration
- **MCP Support**: Basic Model Context Protocol
- **Project detection**: Auto-detects project types

#### Skill Sharing Methods:
```
~/.opencode/
├── config.json         # Global configuration
└── mcp/                # MCP server definitions
```

#### Limitations:
- No formal skill/toolkit system
- Skills are configuration-based, not programmatic
- Limited community skill ecosystem
- Basic MCP implementation

---

### 5. Cursor CLI

**Shared Skills Support: ❌ MINIMAL**

Cursor CLI has the least support for shared skills:

#### Current State:
- **No formal skill system**: Relies on Cursor IDE features
- **Limited MCP**: Partial MCP support through IDE
- **No extension mechanism**: CLI is primarily for headless agent execution
- **Rules files**: Basic `.cursorrules` for project context

#### Skill Sharing Methods:
```
project/
├── .cursorrules        # Basic project rules
└── .cursor/
    └── rules/          # Additional rules files
```

#### Limitations:
- Skill features tied to Cursor IDE, not CLI
- No standalone plugin/toolkit system
- `.cursorrules` is instruction-only, not executable
- MCP servers must be configured in IDE, not CLI

---

## Shared Skills Feature Comparison

### MCP (Model Context Protocol) Support

| Agent | MCP Version | Server Support | Resource Support | Tool Calling |
|-------|-------------|----------------|------------------|--------------|
| Claude Code | 1.0+ | ✅ Full | ✅ Full | ✅ Full |
| Goose | 1.0+ | ✅ Full | ✅ Full | ✅ Full |
| Codex | 1.0 | ✅ Yes | ⚠️ Partial | ✅ Yes |
| OpenCode | 1.0 | ⚠️ Basic | ⚠️ Basic | ⚠️ Basic |
| Cursor CLI | N/A | ⚠️ Via IDE | ⚠️ Via IDE | ⚠️ Via IDE |

### Skill Distribution Methods

| Agent | Local Files | Git Repos | Package Manager | Marketplace |
|-------|-------------|-----------|-----------------|-------------|
| Claude Code | ✅ | ✅ | ✅ npm | ❌ |
| Goose | ✅ | ✅ | ⚠️ | ✅ Community |
| Codex | ✅ | ✅ | ❌ | ❌ |
| OpenCode | ✅ | ⚠️ | ❌ | ❌ |
| Cursor CLI | ✅ | ⚠️ | ❌ | ❌ |

### Skill Types Supported

| Agent | Instructions | Executable Scripts | API Integrations | Custom Tools |
|-------|--------------|-------------------|------------------|--------------|
| Claude Code | ✅ | ✅ Hooks | ✅ MCP | ✅ Commands |
| Goose | ✅ | ✅ Toolkits | ✅ MCP + Native | ✅ Extensions |
| Codex | ✅ | ⚠️ Limited | ✅ MCP | ⚠️ Limited |
| OpenCode | ✅ | ❌ | ⚠️ Basic MCP | ❌ |
| Cursor CLI | ✅ | ❌ | ⚠️ Via IDE | ❌ |

---

## Recommendations

### For Maximum Skill Sharing Capability

**Tier 1 (Full Support):**
1. **Goose** - Best for complex automation with modular toolkits
2. **Claude Code** - Best for MCP ecosystem and developer experience

**Tier 2 (Limited Support):**
3. **Codex** - Adequate for instruction-based skills with MCP
4. **OpenCode** - Basic configuration-based skills

**Tier 3 (Minimal Support):**
5. **Cursor CLI** - Requires IDE for full skill functionality

### Cross-Agent Skill Sharing

The **MCP protocol** is the best path for cross-agent skill sharing:

1. **Write MCP servers** that work with multiple agents
2. **Use standard tool schemas** for compatibility
3. **Package as npm/pip modules** for easy distribution

Example MCP server that works with Claude Code, Goose, and Codex:

```javascript
// mcp-server-example/index.js
const { Server } = require('@modelcontextprotocol/sdk/server');

const server = new Server({
  name: 'shared-skills-server',
  version: '1.0.0'
});

server.addTool({
  name: 'deploy',
  description: 'Deploy the application',
  parameters: {
    environment: { type: 'string', enum: ['dev', 'staging', 'prod'] }
  },
  handler: async ({ environment }) => {
    // Deployment logic
    return { success: true, environment };
  }
});

server.start();
```

---

## Summary Table

| Agent | Skills Support | Best For | Limitation |
|-------|---------------|----------|------------|
| Claude Code | ✅ Full | MCP ecosystem, team sharing | Anthropic-specific features |
| Goose | ✅ Full | Complex automation, toolkits | Steeper learning curve |
| Codex | ⚠️ Limited | Simple instruction-based | No plugin system |
| OpenCode | ⚠️ Limited | Basic project configs | Minimal skill ecosystem |
| Cursor CLI | ❌ Minimal | IDE-coupled workflows | No standalone skills |

---

## Appendix: Skill Configuration Examples

### Claude Code - MCP Server Config
```json
{
  "mcpServers": {
    "github": {
      "command": "npx",
      "args": ["@anthropic/mcp-server-github"]
    },
    "custom": {
      "command": "node",
      "args": ["./mcp-servers/custom.js"]
    }
  }
}
```

### Goose - Toolkit Config
```yaml
toolkits:
  - developer
  - github
  - name: custom-toolkit
    requires:
      - python>=3.10
    tools:
      - name: analyze_code
        command: python scripts/analyze.py
```

### Codex - Instructions
```markdown
# AGENTS.md
## Custom Skills
- analyze: Run static analysis with `npm run lint`
- test: Execute `npm test -- --coverage`
- deploy: Use `./deploy.sh` for deployments
```
