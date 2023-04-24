package controller

import (
	"github.com/remoting/frame/examples/service"
	"github.com/remoting/frame/spring"
)

var beans = spring.New(service.GetSpring())

func init() {
	beans.Reg(NewTestController())
}

func GetSpring() spring.Spring {
	return beans
}
func GetControllers() map[string]any {
	controllers := make(map[string]any)
	controllers["/test"] = spring.GetBean[*TestController](beans)
	return controllers
}
