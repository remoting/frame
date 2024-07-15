package goroutine

import (
	"fmt"
	"github.com/remoting/frame/pkg/logger"
	"time"
)

func Recover(fn ...func(err error)) {
	if p := recover(); p != nil {
		err := fmt.Errorf("%+v\n", p)
		if len(fn) > 0 {
			for _, call := range fn {
				call(err)
			}
		} else {
			logger.Error("error=%s", err.Error())
		}
	}
}

func SafeGo(fn func()) {
	go SafeRun(fn)
}

func SafeRun(fn func()) {
	defer Recover()
	fn()
}
func SafeRetry(fn func()) {
	go retryRun(fn)
}
func retryRun(fn func()) {
	for {
		func() {
			defer Recover(func(err error) {
				logger.Error("Recovered from panic: %v. Retrying in 0.5 seconds...\n", err)
				time.Sleep(500 * time.Millisecond)
			})
			fn()
		}()
	}
}
func Sleep(millisecond int) {
	time.Sleep(time.Duration(millisecond) * time.Millisecond)
}
