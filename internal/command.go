package internal

import (
	"os/exec"
)

func MustRun(command ...string) string {

	f, arr := command[0], command[1:]

	cmd := exec.Command(f, arr...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		PanicError(err)
	}

	return string(output)
}
