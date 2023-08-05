package auth

import "github.com/gin-gonic/gin"

type DefaultAuthService struct {
}

func (_ *DefaultAuthService) Authorization(c *gin.Context) bool {
	user, exists := c.Get("__userInfo__")
	if exists {
		info, ok := user.(UserInfo)
		if ok {
			if info.IsAdmin() {
				return true
			} else {
				return true
			}
		}
	}
	return false
}
