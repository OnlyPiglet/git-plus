package ecmd

import (
	"os"
	"os/exec"
)

func Exec(c string, s ...string) error {
	cmd := exec.Command(c, s...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
