package auth

import (
	"github.com/remoting/frame/server/web"
)

type DefaultAuthService struct {
}

func (_ *DefaultAuthService) Authorization(c *web.Context) bool {
	user, exists := c.Get("__userInfo__")
	if exists {
		info, ok := user.(web.User)
		if ok {
			if info.IsAdministrator() {
				return true
			} else {
				return false
			}
		}
	}
	return false
}
