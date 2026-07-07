package sandbox

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func RunCode(code string, extension string, interpreter string, compiler string) (string, error) {
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
	// for interpreted languages
	if interpreter != "" {

		cmd := exec.CommandContext(ctx, interpreter, tmpFile.Name())

		output, err := cmd.CombinedOutput()

		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("execution timed out after 5 seconds")
		}

		return string(output), err
	}
	// compiled languages
	executable := filepath.Join(os.TempDir(), fmt.Sprintf("%d.exe", time.Now().UnixNano()))

	defer os.Remove(executable)
	//compile
	compilecmd := exec.CommandContext(ctx, compiler, tmpFile.Name(), "-o", executable)

	compileOutput, err := compilecmd.CombinedOutput()

	if err != nil {
		return string(compileOutput), err
	}
	//execute
	runCmd := exec.CommandContext(ctx, executable)

	runOutput, err := runCmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		return "", fmt.Errorf("execution timed out after 5 seconds")
	}

	return string(runOutput), err
}
