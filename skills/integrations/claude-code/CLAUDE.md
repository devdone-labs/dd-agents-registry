# Project Skills (Claude Code)

This project uses shared skills via MCP. The following tools and commands are available.

## Available Skills

### Deploy Skill
Deploy the application to target environments.

**MCP Tool:** `deploy`

**Parameters:**
- `environment` (required): `dev`, `staging`, or `prod`
- `dry_run` (optional): `true` to simulate without changes
- `version` (optional): Specific version to deploy

**Quick Command:**
```
/deploy <environment>
```

**Examples:**
- Deploy to dev: Use deploy tool with environment="dev"
- Dry run to prod: Use deploy tool with environment="prod", dry_run=true

---

### Test Skill
Run the test suite with configurable options.

**MCP Tool:** `test`

**Parameters:**
- `type` (optional): `unit`, `integration`, `e2e`, or `all` (default: all)
- `coverage` (optional): `true` to generate coverage report
- `pattern` (optional): Test file pattern to match

**Quick Command:**
```
/test [type] [--coverage]
```

**Examples:**
- Run all tests: Use test tool
- Run with coverage: Use test tool with coverage=true
- Run unit tests only: Use test tool with type="unit"

---

### Lint Skill
Run code linters and static analysis.

**MCP Tool:** `lint`

**Parameters:**
- `fix` (optional): `true` to auto-fix issues
- `paths` (optional): Array of paths to lint

**Quick Command:**
```
/lint [--fix]
```

**Examples:**
- Check code: Use lint tool
- Fix issues: Use lint tool with fix=true

---

### Format Skill
Format code according to project standards.

**MCP Tool:** `format`

**Parameters:**
- `check` (optional): `true` to check without changes
- `paths` (optional): Array of paths to format

**Quick Command:**
```
/format [--check]
```

**Examples:**
- Format all code: Use format tool
- Check formatting: Use format tool with check=true

---

## Skill Chaining

You can chain skills together for common workflows:

### Pre-commit Check
1. Run lint with fix=true
2. Run format
3. Run test with type="unit"

### Full CI Check
1. Run lint (no fix)
2. Run format with check=true
3. Run test with coverage=true

### Deploy Workflow
1. Run test with type="all"
2. Run lint
3. Deploy with dry_run=true first
4. Deploy to target environment

---

## Project-Specific Notes

Add any project-specific skill configurations or notes here.
