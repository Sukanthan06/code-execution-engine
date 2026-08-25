#!/usr/bin/env bash
set -euo pipefail

echo "Building Docker sandbox image..."
docker build -t code-runner -f executor/docker/Dockerfile executor/docker

echo "Building Go API server binary..."
go build -o bin/server main.go
