package docker

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Sukanthan06/code-execution-engine/config"
	"github.com/Sukanthan06/code-execution-engine/models"
)

func RunCode(code string, extension string, DockerInterpreter string, DockerCompiler string) (*models.ExecutionResult, error) {

	tmpFile, err := os.CreateTemp("", "*"+extension)
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(code); err != nil {
		return nil, err
	}
	tmpFile.Close()

	ctx, cancel := context.WithTimeout(context.Background(), config.AppConfig.Timeout)
	defer cancel()

	absPath, err := filepath.Abs(tmpFile.Name())

	if err != nil {
		return nil, err
	}
	args := []string{
		"run",
		"--rm",
		"--network", config.AppConfig.NetworkMode,
		"--memory=" + config.AppConfig.MemoryLimit,
		"--cpus=" + config.AppConfig.CPULimit,
		"--read-only",
		"--tmpfs", "/tmp:rw,exec",
		"-v", fmt.Sprintf("%s:/app/main%s:ro", absPath, extension),
		config.AppConfig.DockerImage,
	}

	if DockerInterpreter != "" {
		args = append(args, DockerInterpreter, fmt.Sprintf("/app/main%s", extension))
	} else if DockerCompiler != "" {
		command := fmt.Sprintf("%s /app/main%s -o /tmp/main && /tmp/main", DockerCompiler, extension)
		args = append(args, "bash", "-c", command)
	} else {
		return nil, fmt.Errorf("no valid interpreter or compiler configured")
	}

	cmd := exec.CommandContext(ctx, "docker", args...)
	output, err := cmd.CombinedOutput()

	res := &models.ExecutionResult{
		Output: string(output),
	}

	if ctx.Err() == context.DeadlineExceeded {
		return res, fmt.Errorf("execution timed out after %v", config.AppConfig.Timeout)
	}

	return res, err
}
