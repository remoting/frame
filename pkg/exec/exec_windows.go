package exec

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

type OSExecutor struct{}

func (*OSExecutor) OsExec(cmd string) *Result {
	env := os.Environ()
	command := exec.Command("cmd.exe")
	cdir, _ := os.Getwd()
	command.Dir = cdir
	command.SysProcAttr = &syscall.SysProcAttr{CmdLine: fmt.Sprintf(`/c %s`, cmd), HideWindow: true}
	command.Env = env

	result := &Result{}
	var stdoutBuf, stderrBuf bytes.Buffer
	command.Stdout = &stdoutBuf
	command.Stderr = &stderrBuf
	result.Error = command.Run()
	result.Output = stdoutBuf.Bytes()
	result.ErrMsg = stderrBuf.Bytes()
	return result
}
func (*OSExecutor) OsRun(cmd string) error {

	return nil
}
func (*OSExecutor) Command(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}
