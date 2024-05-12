package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/remoting/frame/errors"
	"github.com/remoting/frame/logger"
	"net/http"
	"runtime"
	"runtime/debug"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
func GetAuthHandlerFunc(authentication *Authentication) gin.HandlerFunc {
	if authentication.AuthService == nil {
		authentication.AuthService = &DefaultAuthService{}
	}
	return authentication.Auth()
}
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
						"code": _err.Code,
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

// Login 登录操作
func Login(r http.ResponseWriter, userID string, tokenKey string) (string, error) {
	token, err := authentication.genToken(userID, tokenKey)
	if err != nil {
		return token, errors.NewRestError(10, "Token生成错误")
	} else {
		cookie := http.Cookie{Name: tokenName, Value: token, Path: "/"}
		http.SetCookie(r, &cookie)
		return token, errors.NewRestError(0, "")
	}
}

// LoginByApp 登录操作
func LoginByApp(userID string, tokenKey string) (string, error) {
	token, err := authentication.genToken(userID, tokenKey)
	if err != nil {
		return "", err
	} else {
		return token, nil
	}
}

// Logout 注销操作
func Logout(r *gin.Context) {
	r.SetCookie(tokenName, "", -1, "/", "", false, false)
	token := authentication.getTokenByRequest(r)
	if len(token) > 0 {
		Del(token)
	}
}
