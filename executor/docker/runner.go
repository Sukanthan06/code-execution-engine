package docker

import (
	"os/exec"
)

func TestDocker() (string, error) {
	cmd := exec.Command(
		"docker",
		"run",
		"--rm",
		"hello-world",
	)

	output, err := cmd.CombinedOutput()

	return string(output), err
}
