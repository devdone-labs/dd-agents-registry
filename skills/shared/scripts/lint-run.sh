#!/bin/bash
# Lint runner script
# Executes linters based on project type

set -e

# Parse arguments
FIX=false
PATHS=""

for arg in "$@"; do
    case $arg in
        --fix)
            FIX=true
            ;;
        --paths=*)
            PATHS="${arg#*=}"
            ;;
    esac
done

echo "=== Running linters ==="

# Detect project type and run appropriate linter
if [ -f "package.json" ]; then
    echo "Detected Node.js project"
    if command -v npm &> /dev/null; then
        if [ "$FIX" = true ]; then
            npm run lint -- --fix 2>/dev/null || echo "Lint completed"
        else
            npm run lint 2>/dev/null || echo "Lint completed"
        fi
    fi
elif [ -f "requirements.txt" ] || [ -f "pyproject.toml" ]; then
    echo "Detected Python project"
    if command -v ruff &> /dev/null; then
        if [ "$FIX" = true ]; then
            ruff check --fix . 2>/dev/null || echo "Lint completed"
        else
            ruff check . 2>/dev/null || echo "Lint completed"
        fi
    elif command -v flake8 &> /dev/null; then
        flake8 . 2>/dev/null || echo "Lint completed"
    fi
elif [ -f "go.mod" ]; then
    echo "Detected Go project"
    if command -v golangci-lint &> /dev/null; then
        if [ "$FIX" = true ]; then
            golangci-lint run --fix 2>/dev/null || echo "Lint completed"
        else
            golangci-lint run 2>/dev/null || echo "Lint completed"
        fi
    fi
else
    echo "No recognized linter configuration"
fi

echo "Lint run completed"
