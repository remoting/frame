package exec

import (
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
	exitCode := 0
	output, err := command.CombinedOutput()
	//fmt.Println(string(output))

	if err != nil {
		// 命令执行出错
		exitErr, ok := err.(*exec.ExitError)
		if ok {
			// 获取命令的退出状态
			exitCode = exitErr.ExitCode()
		}
	}
	//output = []byte(ConvertByte2String(output, "GB18030"))

	//fmt.Println(string(output))
	return &Result{
		Code: exitCode,
		Data: output,
		Err:  err,
	}
}
func (*OSExecutor) OsRun(cmd string) error {

	return nil
}
func (*OSExecutor) Command(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}
