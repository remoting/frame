package auth

import (
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/remoting/frame/logger"
)

var tokenName = "jwt-token"
var authentication *Authentication

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
		_token := auth.getTokenByRequest(c)
		_userInfo := auth.getUserByRequest(c, auth.UserService)
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
			if auth.AuthService.Authorization(c) {
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

func (*Authentication) setUserInfoByID(token, userID string, service UserInfoService) UserInfo {
	userInfo, err := service.GetUserInfoById(userID)
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

func (auth *Authentication) verifyToken(tokenString string) (string, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(strings.TrimSpace(auth.TokenSecretPub)))
	if err != nil {
		return "", err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return "", err
	}
	uid, ok := token.Claims.(jwt.MapClaims)["uid"]
	if ok {
		return uid.(string), nil
	}
	return "", err
}
func (*Authentication) genToken(userID string, tokenSecret string) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(strings.TrimSpace(tokenSecret)))
	if err != nil {
		return "", err
	}
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["uid"] = userID
	token.Claims = claims
	return token.SignedString(key)
}

// GetTokenByRequest 获取用户
func (*Authentication) getTokenByRequest(r *gin.Context) string {
	token := ""
	if len(r.Request.Header.Get(tokenName)) > 0 {
		token = r.Request.Header.Get(tokenName)
	} else {
		cook, err := r.Request.Cookie(tokenName)
		if err == nil && len(cook.Value) > 0 {
			token = cook.Value
		}
	}
	return token
}

// GetUserByRequest 获取用户
func (auth *Authentication) getUserByRequest(r *gin.Context, service UserInfoService) UserInfo {
	token := auth.getTokenByRequest(r)
	if len(token) <= 0 {
		return nil
	}
	userID, err := auth.verifyToken(token)
	if err == nil && len(userID) > 0 {
		// 内存缓存里面有就获取出来返回，如果没有就从数据库获取出来放入缓存
		userInfo := Get(token)
		if userInfo != nil {
			//获取一次修改一次访问时间
			TouchUserInfoTime(token)
			return userInfo
		} else {
			// 缓存里面没有，去数据库取一下，然后存入缓存
			return auth.setUserInfoByID(token, userID, service)
		}
	} else {
		r.SetCookie(tokenName, "", -1, "/", "", false, false)
		return nil
	}
}