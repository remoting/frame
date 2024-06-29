package web

import (
	"github.com/remoting/frame/pkg/errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func (router *RouterGroup) AddRouter(prefix string, controller interface{}) {
	t := reflect.TypeOf(controller)
	kind := t.Kind()
	// 结构体指针
	if kind == reflect.Ptr {
		routes := router.Group(prefix)
		for i := 0; i < t.NumMethod(); i++ {
			m := t.Method(i)
			// 判断是否为控制器方法
			if m.Type.NumIn() == 2 && m.Type.In(1) == reflect.TypeOf(&Context{}) {
				if strings.HasPrefix(m.Name, "Get") {
					routes.GET(m.Name, proxyHandlerFunc(m, controller))
				} else {
					routes.POST(m.Name, proxyHandlerFunc(m, controller))
				}
			}
		}
	}
}
func proxyHandlerFunc(method reflect.Method, object interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := &Context{
			Context: c,
		}
		params := []reflect.Value{reflect.ValueOf(object), reflect.ValueOf(context)}
		result := method.Func.Call(params)
		if method.Type.NumOut() > 0 && len(result) > 0 {
			// 判断最后一个返回值是不是 error
			err := result[len(result)-1].Interface()
			if err != nil {
				typ := reflect.TypeOf(err).String()
				if typ == "*errors.errorString" {
					c.JSON(http.StatusOK, gin.H{
						"code": "500",
						"msg":  err.(error).Error(),
						"data": nil,
					})
				} else if typ == "errors.RestError" || typ == "*errors.RestError" {
					c.JSON(http.StatusOK, gin.H{
						"code": err.(errors.RestError).Code,
						"msg":  err.(errors.RestError).Error(),
						"data": nil,
					})
				} else {
					resultProxyHandlerFunc(result, context)
				}
			} else {
				resultProxyHandlerFunc(result, context)
			}
		}
	}
}

func resultProxyHandlerFunc(result []reflect.Value, context *Context) {
	// 判断返回值类型
	obj := result[0].Interface()
	if obj == nil {
		context.JSON(http.StatusOK, &Result{
			Code: 0, Message: "", Data: nil,
		})
	} else {
		if reflect.TypeOf(obj) == reflect.TypeOf(&Result{}) || reflect.TypeOf(obj) == reflect.TypeOf(Result{}) {
			context.JSON(http.StatusOK, obj)
		} else {
			context.JSON(http.StatusOK, &Result{
				Code: 0, Message: "", Data: obj,
			})
		}
	}
}
