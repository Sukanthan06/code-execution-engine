package sandbox

import (
	"os"
	"os/exec"
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

	cmd := exec.Command("python", tmpFile.Name())

	output, err := cmd.CombinedOutput()

	return string(output), err

}
