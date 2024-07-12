package tools

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/remoting/frame/pkg/logger"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("[%s] [ACCESS] [%s] %s %s %d %s %s %s %s %s\n",
			logger.Conf.Prefix,
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			param.ClientIP,
			param.Request.Proto,
			param.StatusCode,
			param.Method,
			param.Latency,
			param.Request.URL,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}
