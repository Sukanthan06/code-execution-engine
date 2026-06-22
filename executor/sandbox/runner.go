package sandbox

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func RunPython(code string) (string, error) {
	tmpFile, err := os.CreateTemp("", "*.py")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(code); err != nil {
		return "", err
	}

	tmpFile.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "python", tmpFile.Name())

	output, err := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		return "", fmt.Errorf("execution timed out after 5 seconds")
	}

	return string(output), err

}
