# Code Execution Engine

[![CI Pipeline](https://github.com/Sukanthan06/code-execution-engine/actions/workflows/ci.yml/badge.svg)](https://github.com/Sukanthan06/code-execution-engine/actions/workflows/ci.yml)

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

## Continuous Integration (GitHub Actions)

Every `push` and `pull_request` to `main` / `master` automatically triggers the GitHub Actions CI pipeline (`.github/workflows/ci.yml`):

1. **Go Verification**:
   - Formats check (`gofmt`).
   - Static analysis (`go vet ./...`).
   - Unit tests (`go test ./...`).
   - Compilation check (`go build ./...`).
2. **Docker Build**:
   - Builds the application Docker sandbox image (`code-runner`).
3. **Real Integration Tests (`scripts/ci-test.sh`)**:
   - Launches the live Go API server.
   - Waits for `/health` endpoint readiness.
   - Executes real integration tests for **Python**, **JavaScript**, **C**, and **C++**.
   - Validates **timeout enforcement**, **C/C++ compilation errors**, and **network isolation**.
   - Verifies returned JSON `output`, `status`, and `exit_code`.

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
