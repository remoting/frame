package tools

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/remoting/frame/pkg/errors"
	"github.com/remoting/frame/pkg/logger"
	"net/http"
	"reflect"
	"runtime"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				//打印错误堆栈信息
				if reflect.TypeOf(r).String() == "*errors.fundamental" {
					// 错误里面包含了堆栈
					logger.ErrorSkip("%+v\n", r)
				} else {
					logger.Error("%s\n", fmt.Sprintf("%+v", r))
				}
				switch _err := r.(type) {
				case errors.RestError:
					c.JSON(http.StatusOK, gin.H{
						"code": _err.Code(),
						"msg":  _err.Error(),
						"data": nil,
					})
				case runtime.Error:
					c.JSON(http.StatusOK, gin.H{
						"code": "500",
						"msg":  _err.Error(),
						"data": nil,
					})
				case error:
					c.JSON(http.StatusOK, gin.H{
						"code": 500,
						"msg":  _err.Error(),
						"data": nil,
					})
				default:
					c.JSON(http.StatusOK, gin.H{
						"code": "500",
						"msg":  "unknown",
						"data": nil,
					})
				}
				c.Abort()
			}
		}()
		//加载完 defer recover，继续后续接口调用
		c.Next()
	}
}
