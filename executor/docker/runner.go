package docker

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/Sukanthan06/code-execution-engine/config"
	"github.com/Sukanthan06/code-execution-engine/models"
)

const (
	dockerCmd            = "docker"
	containerAppDir      = "/app/main"
	containerTmpBin      = "/tmp/main"
	compileErrorExitCode = 100
)

func RunCode(code string, extension string, dockerInterpreter string, dockerCompiler string) (*models.ExecutionResult, error) {

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

	containerSourcePath := containerAppDir + extension

	args := []string{
		"run",
		"--rm",
		"--network", config.AppConfig.NetworkMode,
		"--memory=" + config.AppConfig.MemoryLimit,
		"--cpus=" + config.AppConfig.CPULimit,
		"--read-only",
		"--tmpfs", "/tmp:rw,exec",
		"-v", fmt.Sprintf("%s:%s:ro", absPath, containerSourcePath),
		config.AppConfig.DockerImage,
	}

	if dockerInterpreter != "" {
		args = append(args, dockerInterpreter, containerSourcePath)
	} else if dockerCompiler != "" {
		command := fmt.Sprintf("%s %s -o %s || exit %d; %s", dockerCompiler, containerSourcePath, containerTmpBin, compileErrorExitCode, containerTmpBin)
		args = append(args, "bash", "-c", command)
	} else {
		return nil, fmt.Errorf("no valid interpreter or compiler configured")
	}

	cmd := exec.CommandContext(ctx, dockerCmd, args...)
	startTime := time.Now()
	output, err := cmd.CombinedOutput()
	duration := time.Since(startTime).Milliseconds()

	res := &models.ExecutionResult{
		Output:    string(output),
		RuntimeMS: duration,
	}

	if ctx.Err() == context.DeadlineExceeded {
		res.Status = models.StatusTimeLimitExceeded
		res.Error = fmt.Sprintf("execution timed out after %v", config.AppConfig.Timeout)
		return res, fmt.Errorf("execution timed out after %v", config.AppConfig.Timeout)
	}

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			res.ExitCode = exitErr.ExitCode()
			if dockerCompiler != "" && res.ExitCode == compileErrorExitCode {
				res.Status = models.StatusCompilationError
				res.Error = "compilation error"
			} else {
				res.Status = models.StatusRuntimeError
				res.Error = "runtime error"
			}
		} else {
			res.Status = models.StatusInternalError
			res.Error = err.Error()
		}
		return res, err
	}

	res.Status = models.StatusSuccess
	res.ExitCode = 0
	return res, nil
}
