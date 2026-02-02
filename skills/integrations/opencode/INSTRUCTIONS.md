# Shared Skills Instructions (OpenCode)

This project uses shared skills. Follow these instructions to execute skills.

## Available Skills

### Deploy

Deploy the application to a target environment.

**MCP Tool:** `deploy` (use if available)

**Manual Execution:**
```bash
# Step 1: Validate
./skills/shared/scripts/deploy-validate.sh <environment>

# Step 2: Build
./skills/shared/scripts/deploy-build.sh --env=<environment>

# Step 3: Deploy
./skills/shared/scripts/deploy-execute.sh --env=<environment>
```

**Environments:** dev, staging, prod

---

### Test

Run the test suite.

**MCP Tool:** `test` (use if available)

**Manual Execution:**
```bash
# Run all tests
./skills/shared/scripts/test-run.sh --type=all

# With coverage
./skills/shared/scripts/test-run.sh --type=all --coverage

# Specific type
./skills/shared/scripts/test-run.sh --type=unit
```

---

### Lint

Run code linters.

**MCP Tool:** `lint` (use if available)

**Manual Execution:**
```bash
# Check only
./skills/shared/scripts/lint-run.sh

# Auto-fix
./skills/shared/scripts/lint-run.sh --fix
```

---

### Format

Format code.

**MCP Tool:** `format` (use if available)

**Manual Execution:**
```bash
# Format all
./skills/shared/scripts/format-run.sh

# Check only
./skills/shared/scripts/format-run.sh --check
```

---

## Notes

- Prefer MCP tools when available for better integration
- Fall back to manual scripts if MCP fails
- Scripts auto-detect project type (Node, Python, Go, etc.)
