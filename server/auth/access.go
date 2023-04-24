package auth

import "github.com/gin-gonic/gin"

func authorization(c *gin.Context) bool {
	user, exists := c.Get("__userInfo__")
	if exists {
		info, ok := user.(UserInfo)
		if ok {
			if info.IsAdmin() {
				return true
			}
		}
	}
	return false
}
