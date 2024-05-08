package web

import (
	"github.com/gin-gonic/gin"
	"github.com/remoting/frame/errors"
	"github.com/remoting/frame/json"
	"github.com/remoting/frame/reflect"
	"github.com/remoting/frame/server/auth"
)

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    any    `json:"data"`
}
type Context struct {
	*gin.Context
}

func (c *Context) BindOBJ(t any) {
	err := c.BindJSON(t)
	c.CheckError(err)
}
func (c *Context) ParseJSON() *json.Object {
	obj := &json.Object{}
	err := c.ShouldBindJSON(obj)
	c.CheckError(err)
	return obj
}
func (c *Context) CheckError(exception any) {
	if !reflect.IsNil(exception) {
		panic(exception)
	}
}
func (c *Context) GetUserInfo() auth.UserInfo {
	user, _ := c.Get("__userInfo__")
	info, ok := user.(auth.UserInfo)
	if ok {
		return info
	}
	return nil
}
func (c *Context) Fail(code int, message string) {
	//c.JSON(200, &Result{Code: code, Message: message})
	panic(errors.NewRestError(code, message))
}

func (c *Context) Success(data any) {
	c.JSON(200, &Result{Code: 0, Message: "", Data: data})
}
