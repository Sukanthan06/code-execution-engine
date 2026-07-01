package sandbox

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func RunCode(code string, command string, extension string) (string, error) {
	tmpFile, err := os.CreateTemp("", "*"+extension)
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

	cmd := exec.CommandContext(ctx, command, tmpFile.Name())

	output, err := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		return "", fmt.Errorf("execution timed out after 5 seconds")
	}

	return string(output), err

}
