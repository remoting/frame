package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/remoting/frame/pkg/errors"
	"net/http"
	"strings"
	"time"
)

// Login 登录操作
func Login(r http.ResponseWriter, userID string, tokenKey string, expires time.Duration) (string, errors.RestError) {
	token, err := GenToken(userID, tokenKey,expires)
	if err != nil {
		return token, errors.NewRestError(10, "Token生成错误")
	} else {
		cookie := http.Cookie{Name: tokenName, Value: token, Path: "/"}
		http.SetCookie(r, &cookie)
		return token, errors.NewRestError(0, "")
	}
}

// LoginByApp 登录操作
func LoginByApp(userID string, tokenKey string, expires time.Duration) (string, error) {
	token, err := GenToken(userID, tokenKey,  expires )
	if err != nil {
		return "", err
	} else {
		return token, nil
	}
}

// Logout 注销操作
func Logout(r *gin.Context) {
	r.SetCookie(tokenName, "", -1, "/", "", false, false)
	token := GetTokenByRequest(r)
	if len(token) > 0 {
		Del(token)
	}
}

func GetTokenByRequest(r *gin.Context) string {
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
func GenToken(userID string, tokenSecret string, expires time.Duration) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(strings.TrimSpace(tokenSecret)))
	if err != nil {
		return "", err
	}
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	//time.Hour * time.Duration(24)
	claims["exp"] = time.Now().Add(expires).Unix()
	claims["iat"] = time.Now().Unix()
	claims["uid"] = userID
	token.Claims = claims
	return token.SignedString(key)
}
func VerifyToken(tokenSecret, tokenString string) (string, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(strings.TrimSpace(tokenSecret)))
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
