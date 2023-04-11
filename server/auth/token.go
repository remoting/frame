package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/remoting/frame/errors"
	"github.com/remoting/frame/logger"
)

// UserCache 缓存
type UserCache struct {
	LifeCircle int64
	UserInfo   UserInfo
	TouchTime  int64
}

/*
type UserInfo struct {
	UserID   string `json:"userId"`   // 用户ID
	UserName string `json:"userName"` // 登录名
}
*/

// TokenCache 缓存用户与token的映射关系，后期需要定期清理功能
var TokenCache map[string]*UserCache = make(map[string]*UserCache, 0)

func setUserInfoByID(token, userID string, service UserInfoService) UserInfo {
	userInfo, err := service.GetUserInfoById(userID)
	if err != nil {
		logger.Warn(").Error(", err.Error())
		return nil
	} else {
		TokenCache[token] = &UserCache{
			UserInfo:   userInfo,
			LifeCircle: 10 * int64(time.Minute),
			TouchTime:  time.Now().UnixNano() / int64(time.Millisecond),
		}
		return userInfo
	}
}
func touchUserInfoTime(token string) {
	TokenCache[token].TouchTime = time.Now().UnixNano() / int64(time.Millisecond)
}

// getUserByToken 获取用户
func getUserByToken(token string) UserInfo {
	cache, ok := TokenCache[token]
	if ok {
		if cache.LifeCircle <= 0 {
			return cache.UserInfo
		}
		ctime := time.Now().UnixNano() / int64(time.Millisecond)
		//换成失效判断
		if ctime-cache.TouchTime > cache.LifeCircle {
			delete(TokenCache, token)
			return nil
		}
		cache.TouchTime = ctime
		return cache.UserInfo
	}
	return nil
}
func verifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
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
func genToken(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["uid"] = userID
	token.Claims = claims
	return token.SignedString([]byte(tokenSecret))
}

// GetTokenByRequest 获取用户
func getTokenByRequest(r *gin.Context) string {
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
func getUserByRequest(r *gin.Context, service UserInfoService) UserInfo {
	token := getTokenByRequest(r)
	if len(token) <= 0 {
		return nil
	}
	userID, err := verifyToken(token)
	if err == nil && len(userID) > 0 {
		// 内存缓存里面有就获取出来返回，如果没有就从数据库获取出来放入缓存
		userInfo := getUserByToken(token)
		if userInfo != nil {
			//获取一次修改一次访问时间
			touchUserInfoTime(token)
			return userInfo
		} else {
			// 缓存里面没有，去数据库取一下，然后存入缓存
			return setUserInfoByID(token, userID, service)
		}
	} else {
		r.SetCookie(tokenName, "", -1, "/", "", false, false)
		return nil
	}
}

// Login 登录操作
func Login(r http.ResponseWriter, userID string) errors.RestError {
	token, err := genToken(userID)
	if err != nil {
		return errors.New(10, "Token生成错误")
	} else {
		cookie := http.Cookie{Name: tokenName, Value: token, Path: "/"}
		http.SetCookie(r, &cookie)
		return errors.New(0, "")
	}
}

// LoginByApp 登录操作
func LoginByApp(userID string) (string, error) {
	token, err := genToken(userID)
	if err != nil {
		return "", err
	} else {
		return token, nil
	}
}

// Logout 注销操作
func Logout(r *gin.Context) {
	r.SetCookie(tokenName, "", -1, "/", "", false, false)
	token := getTokenByRequest(r)
	if len(token) > 0 {
		delete(TokenCache, token)
	}
}
