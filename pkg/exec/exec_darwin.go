package exec

import (
	"os"
	"os/exec"
)

type OSExecutor struct{}

func (*OSExecutor) OsExec(cmd string) *Result {
	env := os.Environ()
	command := exec.Command("bash", "-c", cmd)
	command.Env = env
	exitCode := 0
	output, err := command.CombinedOutput()

	if err != nil {
		// 命令执行出错
		exitErr, ok := err.(*exec.ExitError)
		if ok {
			// 获取命令的退出状态
			exitCode = exitErr.ExitCode()
		}
	}
	return &Result{
		Code: exitCode,
		Data: output,
		Err:  err,
	}
}

func (*OSExecutor) Command(name string, args ...string) *exec.Cmd {
	return exec.Command(name, args...)
}
