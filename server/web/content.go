package web

import (
	"github.com/gin-gonic/gin"
	"github.com/remoting/frame/pkg/errors"
	"github.com/remoting/frame/pkg/json"
	"github.com/remoting/frame/pkg/reflect"
)

type HandlerFunc func(*Context)
type Dict map[string]any
type List []any
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    any    `json:"data"`
}
type Context struct {
	*gin.Context
}

func (c *Context) ClientIP() string {
	if c.GetHeader("RemoteIp") != "" {
		return c.GetHeader("RemoteIp")
	}
	if c.GetHeader("X-Real-IP") != "" {
		return c.GetHeader("X-Real-IP")
	}
	return c.Context.ClientIP()
}
func (c *Context) BindOBJ(t any) {
	err := c.BindJSON(t)
	c.CheckError(err)
}
func (c *Context) ParseJSON() json.Object {
	obj := json.Object{}
	err := c.ShouldBindJSON(&obj)
	c.CheckError(err)
	return obj
}
func (c *Context) CheckError(exception any) {
	if !reflect.IsNil(exception) {
		panic(exception)
	}
}
func (c *Context) GetUserInfo() User {
	user, _ := c.Get("__userInfo__")
	info, ok := user.(User)
	if ok {
		return info
	}
	return nil
}
func (c *Context) Fail(code int, message string) {
	//c.JSON(200, &Result{Code: code, Message: message})
	panic(errors.NewRestError(code, message))
}
func (c *Context) Error(err errors.RestError) {
	c.JSON(200, &Result{Code: err.Code(), Message: err.Error(), Data: nil})
	c.Abort()
}
func (c *Context) Success(data any) {
	c.JSON(200, &Result{Code: 0, Message: "", Data: data})
}
