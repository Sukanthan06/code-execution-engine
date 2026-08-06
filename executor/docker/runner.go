package docker

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func RunCode(code string, extension string, DockerInterpreter string, DockerCompiler string) (string, error) {

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

	absPath, err := filepath.Abs(tmpFile.Name())

	if err != nil {
		return "", err
	}

	if DockerInterpreter != "" {
		cmd := exec.CommandContext(
			ctx,
			"docker",
			"run",
			"--rm",
			"--network",
			"none",
			"--memory=128m",
			"--cpus=1",
			"--read-only",
			"--tmpfs",
			"/tmp:rw,exec",
			"-v",
			fmt.Sprintf("%s:/app/main%s", absPath, extension),
			"code-runner",
			DockerInterpreter,
			fmt.Sprintf("/app/main%s", extension),
		)

		output, err := cmd.CombinedOutput()

		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("execution timed out after 5 seconds")
		}

		return string(output), err
	}
	if DockerCompiler != "" {
		command := fmt.Sprintf(
			"%s /app/main%s -o /tmp/main && /tmp/main",
			DockerCompiler,
			extension,
		)
		cmd := exec.CommandContext(
			ctx,
			"docker",
			"run",
			"--rm",
			"--network",
			"none",
			"--memory=128m",
			"--cpus=1",
			"--read-only",
			"--tmpfs",
			"/tmp:rw,exec",
			"-v",
			fmt.Sprintf("%s:/app/main%s", absPath, extension),
			"code-runner",
			"bash",
			"-c",
			command,
		)
		output, err := cmd.CombinedOutput()

		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("execution timed out after 5 seconds")
		}
		return string(output), err
	}
	return "", fmt.Errorf("no valid interpreter or compiler configured")

}
