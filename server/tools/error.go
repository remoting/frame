package tools

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/remoting/frame/pkg/errors"
	"github.com/remoting/frame/pkg/logger"
	"net/http"
	"runtime"
	"runtime/debug"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				//打印错误堆栈信息
				errMsg := fmt.Sprintf("%+v", r)
				logger.Error("%s\n%s", errMsg, string(debug.Stack()))
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
						"msg":  errMsg,
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
