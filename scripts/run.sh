#!/usr/bin/env bash
set -euo pipefail

echo "Starting code-execution-engine API server..."
exec go run main.go
