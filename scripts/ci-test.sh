#!/usr/bin/env bash
set -euo pipefail

BASE_URL="http://localhost:8080"
HEALTH_URL="${BASE_URL}/health"
EXECUTE_URL="${BASE_URL}/execute"

echo "=== 1. Waiting for server to become healthy at ${HEALTH_URL} ==="
MAX_ATTEMPTS=30
ATTEMPT=0

until curl -s -f "${HEALTH_URL}" > /dev/null; do
    ATTEMPT=$((ATTEMPT + 1))
    if [ "${ATTEMPT}" -ge "${MAX_ATTEMPTS}" ]; then
        echo "[ERROR] Server failed to respond to /health within ${MAX_ATTEMPTS} seconds."
        exit 1
    fi
    sleep 1
done

echo "[SUCCESS] Server is healthy!"
echo ""

FAILURES=0

run_test() {
    local name="$1"
    local lang="$2"
    local code="$3"
    local expected_status="$4"
    local expected_output_substring="${5:-}"

    echo "--- Running test: ${name} (${lang}) ---"

    payload=$(jq -n --arg l "${lang}" --arg c "${code}" '{language: $l, code: $c}')
    response=$(curl -s -X POST "${EXECUTE_URL}" \
        -H "Content-Type: application/json" \
        -d "${payload}")

    actual_status=$(echo "${response}" | jq -r '.status // empty')
    actual_output=$(echo "${response}" | jq -r '.output // empty')
    actual_error=$(echo "${response}" | jq -r '.error // empty')

    echo "Response: ${response}"

    if [ "${actual_status}" != "${expected_status}" ]; then
        echo "[FAIL] ${name}: Expected status '${expected_status}', got '${actual_status}'"
        FAILURES=$((FAILURES + 1))
        return
    fi

    if [ -n "${expected_output_substring}" ]; then
        if [[ "${actual_output}" != *"${expected_output_substring}"* ]]; then
            echo "[FAIL] ${name}: Expected output substring '${expected_output_substring}', got '${actual_output}'"
            FAILURES=$((FAILURES + 1))
            return
        fi
    fi

    echo "[PASS] ${name}"
    echo ""
}

# 1. Python Execution (Success)
run_test "Python Execution" "python" "print('Hello Python')" "SUCCESS" "Hello Python"

# 2. JavaScript Execution (Success)
run_test "JavaScript Execution" "javascript" "console.log('Hello JavaScript');" "SUCCESS" "Hello JavaScript"

# 3. C Execution (Success)
c_code=$'#include <stdio.h>\nint main() { printf("Hello C\\n"); return 0; }'
run_test "C Execution" "c" "${c_code}" "SUCCESS" "Hello C"

# 4. C++ Execution (Success)
cpp_code=$'#include <iostream>\nint main() { std::cout << "Hello C++" << std::endl; return 0; }'
run_test "C++ Execution" "cpp" "${cpp_code}" "SUCCESS" "Hello C++"

# 5. Python Timeout
run_test "Python Timeout" "python" "while True: pass" "TIME_LIMIT_EXCEEDED"

# 6. C Compilation Error
bad_c_code=$'#include <stdio.h>\nint main() { printf("Missing semicolon") return 0; }'
run_test "C Compilation Error" "c" "${bad_c_code}" "COMPILATION_ERROR"

# 7. C++ Compilation Error
bad_cpp_code=$'#include <iostream>\nint main() { std::cout << "Missing semicolon" return 0; }'
run_test "C++ Compilation Error" "cpp" "${bad_cpp_code}" "COMPILATION_ERROR"

# 8. Security Verification: Network Access Blocked
net_code=$'import socket\ntry:\n    socket.create_connection(("1.1.1.1", 80), timeout=2)\n    print("NETWORK_CONNECTED")\nexcept Exception:\n    print("NETWORK_BLOCKED")'
run_test "Network Isolation Test" "python" "${net_code}" "SUCCESS" "NETWORK_BLOCKED"

# 9. Security Verification: Read-Only Filesystem & Writable /tmp
readonly_code=$'import os\ntry:\n    with open("/app/readonly_test.txt", "w") as f:\n        f.write("fail")\n    print("WRITE_ALLOWED")\nexcept Exception:\n    with open("/tmp/writable_test.txt", "w") as f:\n        f.write("ok")\n    print("READONLY_ENFORCED")'
run_test "Read-Only Filesystem Test" "python" "${readonly_code}" "SUCCESS" "READONLY_ENFORCED"

echo "========================================="
if [ "${FAILURES}" -gt 0 ]; then
    echo "[FAIL] ${FAILURES} integration test(s) failed."
    exit 1
else
    echo "[SUCCESS] All integration tests passed!"
    exit 0
fi
