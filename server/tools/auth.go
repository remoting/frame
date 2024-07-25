package tools

import (
	"github.com/gin-gonic/gin"
	"github.com/remoting/frame/server/auth"
)

func InitAuthentication(config auth.AuthConfig) *auth.Authentication {
	return &auth.Authentication{
		AuthConfig: config,
	}
}

func NewAuthentication(userService auth.UserService, secret string, anonymousPath []string, _authorization auth.Authorization) *auth.Authentication {
	var authorization auth.Authorization
	if _authorization == nil {
		authorization = &auth.DefaultAuthorization{}
	} else {
		authorization = _authorization
	}
	return &auth.Authentication{
		AuthConfig: auth.AuthConfig{
			GetUserService: func() auth.UserService {
				return userService
			},
			GetTokenSecret: func() string {
				return secret
			},
			GetAnonymousPath: func() []string {
				return anonymousPath
			},
			GetAuthService: func() auth.Authorization {
				return authorization
			},
		},
	}
}
func CreateAuthentication(userService auth.UserService, secret string, anonymousPath []string, _authorization auth.Authorization) gin.HandlerFunc {
	return NewAuthentication(userService, secret, anonymousPath, _authorization).HandlerFunc()
}
