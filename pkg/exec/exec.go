package exec

import (
	"bytes"
	"os/exec"
)

type Result struct {
	Output bytes.Buffer
	ErrMsg bytes.Buffer
	Error  error
}
type Executor interface {
	OsExec(cmd string) *Result
	OsRun(cmd string) error
	Command(name string, args ...string) *exec.Cmd
}

var executor Executor

func OsExec(cmd string) *Result {
	if executor == nil {
		executor = &OSExecutor{}
	}
	return executor.OsExec(cmd)
}
func OsRun(cmd string) error {
	if executor == nil {
		executor = &OSExecutor{}
	}
	return executor.OsRun(cmd)
}
func OsCommand(name string, args ...string) *exec.Cmd {
	if executor == nil {
		executor = &OSExecutor{}
	}
	return executor.Command(name, args...)
}
