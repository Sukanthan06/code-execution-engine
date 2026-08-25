# Code Execution Engine

[![Fast CI](https://github.com/Sukanthan06/code-execution-engine/actions/workflows/fast-ci.yml/badge.svg)](https://github.com/Sukanthan06/code-execution-engine/actions/workflows/fast-ci.yml)
[![Docker Integration](https://github.com/Sukanthan06/code-execution-engine/actions/workflows/docker-integration.yml/badge.svg)](https://github.com/Sukanthan06/code-execution-engine/actions/workflows/docker-integration.yml)

A secure, high-performance, multi-language code execution engine built in Go inspired by Judge0 and Piston.

## Overview

The `code-execution-engine` receives untrusted user-submitted code over a Gin REST API and executes it inside isolated, restricted Docker containers with strict security, memory, CPU, and timeout enforcement.

### Supported Languages
- **Python** (via `python3`)
- **JavaScript** (via `node`)
- **C** (via `gcc`)
- **C++** (via `g++`)

---

## Architecture & Docker Sandbox Security

All submissions for supported languages run in isolated Docker containers configured with:
- **`--rm`**: Automatic container removal upon exit.
- **`--network none`**: Strict network isolation preventing external internet access.
- **`--memory 128m`**: 128 MB RAM memory limit.
- **`--cpus 1`**: Enforced 1.0 CPU core limit.
- **`--read-only`**: Container root filesystem mounted read-only.
- **`--tmpfs /tmp:rw,exec`**: `/tmp` mounted as in-memory writable filesystem with execution permission (used by C/C++ compilers).
- **Non-root User**: Container executes under unprivileged `runner` user.
- **Execution Timeout**: 5-second deadline context limit.

---

## REST API Interface

### 1. Health Check
`GET /health`

**Response (`200 OK`)**:
```json
{
  "service": "code-execution-engine",
  "status": "ok"
}
```

### 2. Execute Code
`POST /execute`

**Request Body**:
```json
{
  "language": "python",
  "code": "print('Hello World')"
}
```

**Response (`200 OK`)**:
```json
{
  "output": "Hello World\n",
  "status": "SUCCESS",
  "exit_code": 0,
  "runtime_ms": 142
}
```

**Error Response Statuses**:
- `COMPILATION_ERROR`: C/C++ compilation error.
- `RUNTIME_ERROR`: Process non-zero exit or exception.
- `TIME_LIMIT_EXCEEDED`: Code execution exceeded timeout.
- `UNSUPPORTED_LANGUAGE`: Unregistered language requested.
- `INTERNAL_ERROR`: Execution engine error.

---

## Two-Tier CI Architecture (GitHub Actions)

To minimize GitHub Actions minutes on private repositories while preserving strict quality gates, the CI pipeline is split into two specialized workflows:

### 1. Fast CI (`.github/workflows/fast-ci.yml`)
- **Triggers**: Every `push` to any branch, and every `pull_request`.
- **Purpose**: Rapidly validates Go code formatting (`gofmt`), static analysis (`go vet ./...`), unit tests (`go test ./...`), and compilation (`go build ./...`) without incurring expensive Docker image build time.
- **Concurrency**: Automatically cancels outdated in-progress runs on feature branch pushes.

### 2. Docker Integration CI (`.github/workflows/docker-integration.yml`)
- **Triggers**:
  - Open, updated, or reopened **Pull Requests** targeting `main`/`master`.
  - Pushes directly to `main`/`master`.
  - Manual triggers via `workflow_dispatch`.
- **Purpose**: Runs end-to-end integration tests (`scripts/ci-test.sh`) against a live Go API server and Docker sandbox container.
- **Verification Coverage**:
  - Python, JavaScript, C, and C++ execution correctness.
  - C/C++ compilation error detection.
  - 5-second timeout enforcement (`TIME_LIMIT_EXCEEDED`).
  - Network isolation (`--network none` blocking socket connections).
  - Read-only root filesystem enforcement and writable `/tmp`.

---

## Running Locally

### Build Docker Image
```bash
./scripts/build.sh
```

### Start Server
```bash
./scripts/run.sh
```

### Run Integration Tests
```bash
./scripts/ci-test.sh
```
