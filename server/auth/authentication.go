package auth

import (
	"github.com/remoting/frame/pkg/logger"
	"github.com/remoting/frame/server/web"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	GetUserById(id string) (web.User, error)
}

// Authorization  授权服务
type Authorization interface {
	Authorization(c *web.Context) bool
}

var tokenName = "jwt-token"

type AuthConfig struct {
	GetAuthService   func() Authorization //授权服务
	GetUserService   func() UserService   //用户信息服务
	GetTokenSecret   func() string        //jwt Token 验证公钥
	GetAnonymousPath func() []string      //匿名可访问路径
}
type Authentication struct {
	AuthConfig
}

func (auth *Authentication) isAnonymousPath(path string) bool {
	if path == "/" {
		return true
	}
	for _, v := range auth.GetAnonymousPath() {
		isMatch, _ := regexp.MatchString(v, path)
		if isMatch {
			return true
		}
	}
	return false
}

func (auth *Authentication) HandlerFunc() gin.HandlerFunc {
	var authorization Authorization
	if auth.GetAuthService == nil || auth.GetAuthService() == nil {
		authorization = &DefaultAuthorization{}
	} else {
		authorization = auth.GetAuthService()
	}
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
			if authorization.Authorization(context) {
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
	userInfo, err := auth.GetUserService().GetUserById(userID)
	if err != nil {
		logger.Warn("Error,%s", err.Error())
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
	userID, err := VerifyToken(auth.GetTokenSecret(), token)
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
