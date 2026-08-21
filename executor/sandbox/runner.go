package sandbox

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/Sukanthan06/code-execution-engine/models"
)

func RunCode(code string, extension string, LocalInterpreter string, LocalCompiler string) (*models.ExecutionResult, error) {
	tmpFile, err := os.CreateTemp("", "*"+extension)
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(code); err != nil {
		return nil, err
	}
	tmpFile.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// for interpreted languages
	if LocalInterpreter != "" {

		cmd := exec.CommandContext(ctx, LocalInterpreter, tmpFile.Name())

		startTime := time.Now()
		output, err := cmd.CombinedOutput()
		duration := time.Since(startTime).Milliseconds()

		res := &models.ExecutionResult{
			Output:    string(output),
			RuntimeMS: duration,
		}

		if ctx.Err() == context.DeadlineExceeded {
			res.Status = models.StatusTimeLimitExceeded
			res.Error = "execution timed out after 5 seconds"
			return res, fmt.Errorf("execution timed out after 5 seconds")
		}

		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				res.ExitCode = exitErr.ExitCode()
				res.Status = models.StatusRuntimeError
				res.Error = "runtime error"
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
	// compiled languages
	executable := filepath.Join(os.TempDir(), fmt.Sprintf("%d.exe", time.Now().UnixNano()))

	defer os.Remove(executable)
	//compile
	compilecmd := exec.CommandContext(ctx, LocalCompiler, tmpFile.Name(), "-o", executable)

	compileOutput, err := compilecmd.CombinedOutput()

	if err != nil {
		res := &models.ExecutionResult{
			Output: string(compileOutput),
			Status: models.StatusCompilationError,
			Error:  "compilation error",
		}
		if exitErr, ok := err.(*exec.ExitError); ok {
			res.ExitCode = exitErr.ExitCode()
		}
		return res, err
	}
	//execute
	runCmd := exec.CommandContext(ctx, executable)

	startTime := time.Now()
	runOutput, err := runCmd.CombinedOutput()
	duration := time.Since(startTime).Milliseconds()

	res := &models.ExecutionResult{
		Output:    string(runOutput),
		RuntimeMS: duration,
	}

	if ctx.Err() == context.DeadlineExceeded {
		res.Status = models.StatusTimeLimitExceeded
		res.Error = "execution timed out after 5 seconds"
		return res, fmt.Errorf("execution timed out after 5 seconds")
	}

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			res.ExitCode = exitErr.ExitCode()
			res.Status = models.StatusRuntimeError
			res.Error = "runtime error"
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
