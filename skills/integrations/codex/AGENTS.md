# Project Skills (Codex)

This project uses shared skills. Use the MCP tools when available, or follow manual instructions below.

## Available Skills

### Deploy Skill

**MCP Tool:** `deploy`

**Manual Steps:**
1. Validate environment configuration
2. Build application artifacts
3. Execute deployment to target

**Parameters:**
- `environment`: dev, staging, or prod (required)
- `dry_run`: true/false (optional, default: false)
- `version`: version tag (optional, default: latest)

**Manual Commands:**
```bash
# Validate
./skills/shared/scripts/deploy-validate.sh <environment>

# Build
./skills/shared/scripts/deploy-build.sh --env=<environment> --version=<version>

# Deploy
./skills/shared/scripts/deploy-execute.sh --env=<environment>
```

---

### Test Skill

**MCP Tool:** `test`

**Manual Steps:**
1. Set up test environment
2. Run test suite
3. Generate reports

**Parameters:**
- `type`: unit, integration, e2e, or all (default: all)
- `coverage`: true/false (default: false)
- `pattern`: test file pattern (optional)

**Manual Commands:**
```bash
# Run tests
./skills/shared/scripts/test-run.sh --type=<type> [--coverage] [--pattern=<pattern>]
```

---

### Lint Skill

**MCP Tool:** `lint`

**Manual Steps:**
1. Detect project language
2. Run appropriate linters
3. Report or fix issues

**Parameters:**
- `fix`: true/false (default: false)
- `paths`: specific paths to lint (optional)

**Manual Commands:**
```bash
# Check only
./skills/shared/scripts/lint-run.sh

# Auto-fix
./skills/shared/scripts/lint-run.sh --fix
```

---

### Format Skill

**MCP Tool:** `format`

**Manual Steps:**
1. Detect project formatters
2. Format or check code
3. Report changes

**Parameters:**
- `check`: true/false (default: false)
- `paths`: specific paths to format (optional)

**Manual Commands:**
```bash
# Format code
./skills/shared/scripts/format-run.sh

# Check only
./skills/shared/scripts/format-run.sh --check
```

---

## Common Workflows

### Pre-Commit
1. Run lint with fix=true
2. Run format
3. Run test with type=unit

### CI Pipeline
1. Run lint (no fix)
2. Run format --check
3. Run test with coverage

### Deployment
1. Run all tests
2. Run lint
3. Deploy to target environment
