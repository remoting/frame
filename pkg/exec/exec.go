package exec

import "os/exec"

type Result struct {
	Data []byte
	Code int
	Err  error
}
type Executor interface {
	OsExec(cmd string) *Result
	Command(name string, args ...string) *exec.Cmd
}

var executor Executor

func OsExec(cmd string) *Result {
	if executor == nil {
		executor = &OSExecutor{}
	}
	return executor.OsExec(cmd)
}

func OsCommand(name string, args ...string) *exec.Cmd {
	if executor == nil {
		executor = &OSExecutor{}
	}
	return executor.Command(name, args...)
}
