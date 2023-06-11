package auth

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/remoting/frame/errors"
	"github.com/remoting/frame/json"
)

var tokenName = "jwt-token"
var anonymousPath []string
var tokenSecret = ""

type UserInfo interface {
	UserId() string
	UserName() string
	GetMenus() []*json.Node
	GetRoles() []*json.Dict
	IsAdmin() bool
}
type UserInfoService interface {
	GetUserInfoById(userId string) (UserInfo, error)
}
type AuthConfig struct {
	UserService   UserInfoService
	TokenSecret   string
	AnonymousPath []string
}

func isAnonymousPath(path string) bool {
	if path == "/" {
		return true
	}
	for _, v := range anonymousPath {
		if strings.HasPrefix(path, v) {
			return true
		}
	}
	return false
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
func GetAuthHandlerFunc(config AuthConfig) gin.HandlerFunc {
	return Auth(config.UserService, config.TokenSecret, config.AnonymousPath)
}
func Auth(service UserInfoService, secret string, paths []string) gin.HandlerFunc {
	tokenSecret = secret
	anonymousPath = paths
	return func(c *gin.Context) {
		_token := getTokenByRequest(c)
		_userInfo := getUserByRequest(c, service)
		// 将当前用户信息放在 request 对象上，方便后面的控制器获取当前用户
		c.Set("__userInfo__", _userInfo)
		c.Set("__token__", _token)
		if _userInfo == nil {
			// 未登陆用户判断是否是匿名允许访问的路径
			if isAnonymousPath(c.Request.URL.Path) {
				c.Next()
			} else {
				c.JSON(401, "未认证用户，不能访问")
			}
		} else {
			// 已登陆用户判断是否有当前URL的访问权限
			if authorization(c) {
				c.Next()
			} else {
				c.JSON(403, "未授权用户，不能访问")
			}
		}
	}
}
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				//打印错误堆栈信息
				fmt.Println("error:========", r.(error).Error())
				//debug.PrintStack()
				switch r.(type) {
				case errors.RestError:
					c.JSON(http.StatusOK, gin.H{
						"code": r.(errors.RestError).Code,
						"msg":  r.(errors.RestError).Error(),
						"data": nil,
					})
				case error:
					c.JSON(http.StatusOK, gin.H{
						"code": 500,
						"msg":  r.(error).Error(),
						"data": nil,
					})
				case runtime.Error:
					c.JSON(http.StatusOK, gin.H{
						"code": "500",
						"msg":  r.(error).Error(),
						"data": nil,
					})
				default:
					c.JSON(http.StatusOK, gin.H{
						"code": "500",
						"msg":  "服务器错误",
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
