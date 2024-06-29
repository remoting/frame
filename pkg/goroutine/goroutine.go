package goroutine

import (
	"fmt"
	"github.com/remoting/frame/pkg/logger"
	"runtime/debug"
	"time"
)

func Recover() {
	if p := recover(); p != nil {
		logger.Error("error=%s", fmt.Errorf("%+v\n%s", p, debug.Stack()).Error())
	}
}

func SafeGo(fn func()) {
	go SafeRun(fn)
}

func SafeRun(fn func()) {
	defer Recover()
	fn()
}

func Sleep(millisecond int) {
	time.Sleep(time.Duration(millisecond) * time.Millisecond)
}
