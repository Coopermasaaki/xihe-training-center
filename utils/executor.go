package utils

import (
	"io"
	"os/exec"
)

func RunCmdInput(stdin string, args ...string) (int, error) {
	n := len(args)
	if n == 0 {
		return 0, nil
	}

	cmd := args[0]

	if n > 1 {
		args = args[1:]
	} else {
		args = nil
	}

	c := exec.Command(cmd, args...)
	stdinPipe, err := c.StdinPipe()

	go func() {
		defer stdinPipe.Close()
		_, err = io.WriteString(stdinPipe, stdin)
	}()

	_, err = c.CombinedOutput()

	if err == nil {
		return 0, nil
	}

	return -1, err
}
