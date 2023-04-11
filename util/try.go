package util

import (
	"log"
	"runtime"
)

func TryCatch(try func(), catch ...func(exception error)) {
	defer func() {
		var r any = recover()
		switch r.(type) {
		case runtime.Error:
			log.Println("运行时错误：", r)
		default:
			//catch[0]()
			log.Println("xxx出现问题，已经跳过该问题。。。")
		}
	}()
	try()
}
func Throw(exception interface{}) {
	if exception != nil {
		panic(exception)
	}
}
