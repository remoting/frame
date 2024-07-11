package auth

import (
	"github.com/remoting/frame/server/web"
)

type DefaultAuthService struct {
}

func (_ *DefaultAuthService) Authorization(c *web.Context) bool {
	user, exists := c.Get("__userInfo__")
	if exists {
		_, ok := user.(web.User)
		if ok {
			// 只需要判断是登录状态就表示有权限
			return true
		}
	}
	return false
}
