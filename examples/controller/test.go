package controller

import (
	"github.com/remoting/frame/server/auth"
	"github.com/remoting/frame/server/web"
)

type TestController struct {
}

func NewTestController() *TestController {
	return &TestController{}
}

func (*TestController) Hello(c *web.Context) {

}

func (*TestController) Login(c *web.Context) {
	errs := auth.Login(c.Writer, "test")
	if errs.Code == 0 {
		c.Success("")
	} else {
		c.Fail(errs.Code, errs.Error())
	}
}
