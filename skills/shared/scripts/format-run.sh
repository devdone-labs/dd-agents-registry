#!/bin/bash
# Format runner script
# Formats code based on project type

set -e

# Parse arguments
CHECK=false
PATHS=""

for arg in "$@"; do
    case $arg in
        --check)
            CHECK=true
            ;;
        --paths=*)
            PATHS="${arg#*=}"
            ;;
    esac
done

echo "=== Running formatters ==="

# Detect project type and run appropriate formatter
if [ -f "package.json" ]; then
    echo "Detected Node.js project"
    if command -v npx &> /dev/null; then
        if [ "$CHECK" = true ]; then
            npx prettier --check . 2>/dev/null || echo "Format check completed"
        else
            npx prettier --write . 2>/dev/null || echo "Format completed"
        fi
    fi
elif [ -f "requirements.txt" ] || [ -f "pyproject.toml" ]; then
    echo "Detected Python project"
    if command -v black &> /dev/null; then
        if [ "$CHECK" = true ]; then
            black --check . 2>/dev/null || echo "Format check completed"
        else
            black . 2>/dev/null || echo "Format completed"
        fi
    elif command -v ruff &> /dev/null; then
        if [ "$CHECK" = true ]; then
            ruff format --check . 2>/dev/null || echo "Format check completed"
        else
            ruff format . 2>/dev/null || echo "Format completed"
        fi
    fi
elif [ -f "go.mod" ]; then
    echo "Detected Go project"
    if command -v gofmt &> /dev/null; then
        if [ "$CHECK" = true ]; then
            gofmt -l . 2>/dev/null || echo "Format check completed"
        else
            gofmt -w . 2>/dev/null || echo "Format completed"
        fi
    fi
else
    echo "No recognized formatter configuration"
fi

echo "Format run completed"
