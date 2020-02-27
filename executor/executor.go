package executor

import (
	"os"
	"os/exec"
	"strings"
)

// Executor type to define command to execute
type Executor struct {
	CommandText      string
	CommandDirectory string
	BindOutput       bool
}

//Execute the command
func (e Executor) Execute() error {
	parts := strings.Fields(e.CommandText)
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Dir = e.CommandDirectory
	if e.BindOutput {
		cmd.Stdout = os.Stdout
	}
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}
