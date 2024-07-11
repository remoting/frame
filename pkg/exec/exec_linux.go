package exec

import (
	"bytes"
	"os"
	"os/exec"
)

type OSExecutor struct{}

func (*OSExecutor) OsExec(cmd string) *Result {
	env := os.Environ()
	command := exec.Command("bash", "-c", cmd)
	command.Env = env
	result := &Result{}
	var stdoutBuf, stderrBuf bytes.Buffer
	command.Stdout = &stdoutBuf
	command.Stderr = &stderrBuf
	result.Error = command.Run()
	return result
}

func (*OSExecutor) OsRun(cmd string) error {
	env := os.Environ()
	command := exec.Command("bash", "-c", cmd)
	command.Env = env
	err := command.Run()
	return err
}
func (*OSExecutor) Command(name string, args ...string) *exec.Cmd {
	return exec.Command(name, args...)
}
