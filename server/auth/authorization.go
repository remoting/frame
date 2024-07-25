package auth

import (
	"github.com/remoting/frame/server/web"
)

type DefaultAuthorization struct {
}

func (_ *DefaultAuthorization) Authorization(c *web.Context) bool {
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
