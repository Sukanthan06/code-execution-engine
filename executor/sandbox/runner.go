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

		output, err := cmd.CombinedOutput()
		res := &models.ExecutionResult{
			Output: string(output),
		}

		if ctx.Err() == context.DeadlineExceeded {
			return res, fmt.Errorf("execution timed out after 5 seconds")
		}

		return res, err
	}
	// compiled languages
	executable := filepath.Join(os.TempDir(), fmt.Sprintf("%d.exe", time.Now().UnixNano()))

	defer os.Remove(executable)
	//compile
	compilecmd := exec.CommandContext(ctx, LocalCompiler, tmpFile.Name(), "-o", executable)

	compileOutput, err := compilecmd.CombinedOutput()

	if err != nil {
		return &models.ExecutionResult{Output: string(compileOutput)}, err
	}
	//execute
	runCmd := exec.CommandContext(ctx, executable)

	runOutput, err := runCmd.CombinedOutput()
	res := &models.ExecutionResult{
		Output: string(runOutput),
	}

	if ctx.Err() == context.DeadlineExceeded {
		return res, fmt.Errorf("execution timed out after 5 seconds")
	}

	return res, err
}
