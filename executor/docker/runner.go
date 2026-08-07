package docker

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Sukanthan06/code-execution-engine/config"
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

	ctx, cancel := context.WithTimeout(context.Background(), config.AppConfig.Timeout)
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
			config.AppConfig.NetworkMode,
			"--memory="+config.AppConfig.MemoryLimit,
			"--cpus="+config.AppConfig.CPULimit,
			"--read-only",
			"--tmpfs",
			"/tmp:rw,exec",
			"-v",
			fmt.Sprintf("%s:/app/main%s", absPath, extension),
			config.AppConfig.DockerImage,
			DockerInterpreter,
			fmt.Sprintf("/app/main%s", extension),
		)

		output, err := cmd.CombinedOutput()

		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("execution timed out after %v", config.AppConfig.Timeout)
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
			config.AppConfig.NetworkMode,
			"--memory="+config.AppConfig.MemoryLimit,
			"--cpus="+config.AppConfig.CPULimit,
			"--read-only",
			"--tmpfs",
			"/tmp:rw,exec",
			"-v",
			fmt.Sprintf("%s:/app/main%s", absPath, extension),
			config.AppConfig.DockerImage,
			"bash",
			"-c",
			command,
		)
		output, err := cmd.CombinedOutput()

		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("execution timed out after %v", config.AppConfig.Timeout)
		}
		return string(output), err
	}
	return "", fmt.Errorf("no valid interpreter or compiler configured")

}
