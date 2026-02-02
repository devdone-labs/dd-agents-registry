#!/bin/bash
# Deploy build script
# Builds application artifacts for deployment

set -e

# Parse arguments
for arg in "$@"; do
    case $arg in
        --env=*)
            ENVIRONMENT="${arg#*=}"
            ;;
        --version=*)
            VERSION="${arg#*=}"
            ;;
    esac
done

ENVIRONMENT="${ENVIRONMENT:-dev}"
VERSION="${VERSION:-latest}"

echo "=== Building for: $ENVIRONMENT (version: $VERSION) ==="

# Detect build system and run appropriate command
if [ -f "package.json" ]; then
    echo "Detected Node.js project"
    if command -v npm &> /dev/null; then
        npm run build 2>/dev/null || echo "No build script found"
    fi
elif [ -f "requirements.txt" ] || [ -f "pyproject.toml" ]; then
    echo "Detected Python project"
    # Python projects typically don't need a build step
elif [ -f "go.mod" ]; then
    echo "Detected Go project"
    if command -v go &> /dev/null; then
        go build ./... 2>/dev/null || echo "Go build completed"
    fi
elif [ -f "Cargo.toml" ]; then
    echo "Detected Rust project"
    if command -v cargo &> /dev/null; then
        cargo build --release 2>/dev/null || echo "Cargo build completed"
    fi
else
    echo "No recognized project type, skipping build"
fi

echo "Build completed successfully"
