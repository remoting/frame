package auth

import (
	"github.com/remoting/frame/pkg/logger"
	"github.com/remoting/frame/server/web"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	GetUserById(id string) (web.User, error)
}

// AuthService  授权服务
type AuthService interface {
	Authorization(c *web.Context) bool
}

var tokenName = "jwt-token"

type Authentication struct {
	AuthService   AuthService //授权服务
	UserService   UserService //用户信息服务
	TokenSecret   string      //jwt Token 验证公钥
	AnonymousPath []string    //匿名可访问路径
}

func (auth *Authentication) isAnonymousPath(path string) bool {
	if path == "/" {
		return true
	}
	for _, v := range auth.AnonymousPath {
		if strings.HasPrefix(path, v) {
			return true
		}
	}
	return false
}

func (auth *Authentication) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_token := GetTokenByRequest(c)
		_userInfo := auth.getUserByRequest(c)
		// 将当前用户信息放在 request 对象上，方便后面的控制器获取当前用户
		c.Set("__userInfo__", _userInfo)
		c.Set("__token__", _token)
		if _userInfo == nil {
			// 未登陆用户判断是否是匿名允许访问的路径
			if auth.isAnonymousPath(c.Request.URL.Path) {
				c.Next()
			} else {
				c.JSON(401, map[string]interface{}{
					"code": 401,
					"msg":  "未认证用户，不能访问",
					"data": "",
				})
				c.Abort()
			}
		} else {
			// 已登陆用户判断是否有当前URL的访问权限
			context := &web.Context{
				Context: c,
			}
			if auth.AuthService.Authorization(context) {
				c.Next()
			} else {
				c.JSON(403, map[string]interface{}{
					"code": 403,
					"msg":  "未授权用户，不能访问",
					"data": "",
				})
				c.Abort()
			}
		}
	}
}

func (auth *Authentication) setUserInfoByID(token, userID string) web.User {
	userInfo, err := auth.UserService.GetUserById(userID)
	if err != nil {
		logger.Warn(").Error(", err.Error())
		return nil
	} else {
		Put(token, &UserCache{
			UserInfo:   userInfo,
			LifeCircle: 10 * int64(time.Minute),
			TouchTime:  time.Now().UnixNano() / int64(time.Millisecond),
		})
		return userInfo
	}
}

// GetUserByRequest 获取用户
func (auth *Authentication) getUserByRequest(r *gin.Context) web.User {
	token := GetTokenByRequest(r)
	if len(token) <= 0 {
		return nil
	}
	userID, err := VerifyToken(auth.TokenSecret, token)
	if err == nil && len(userID) > 0 {
		// 内存缓存里面有就获取出来返回，如果没有就从数据库获取出来放入缓存
		userInfo := Get(token)
		if userInfo != nil {
			//获取一次修改一次访问时间
			TouchUserInfoTime(token)
			return userInfo
		} else {
			// 缓存里面没有，去数据库取一下，然后存入缓存
			return auth.setUserInfoByID(token, userID)
		}
	} else {
		r.SetCookie(tokenName, "", -1, "/", "", false, false)
		return nil
	}
}
