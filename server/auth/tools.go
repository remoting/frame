package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/remoting/frame/pkg/errors"
	"net/http"
)

// Login 登录操作
func Login(r http.ResponseWriter, userID string, tokenKey string) (string, errors.RestError) {
	token, err := authentication.genToken(userID, tokenKey)
	if err != nil {
		return token, errors.NewRestError(10, "Token生成错误")
	} else {
		cookie := http.Cookie{Name: tokenName, Value: token, Path: "/"}
		http.SetCookie(r, &cookie)
		return token, errors.NewRestError(0, "")
	}
}

// LoginByApp 登录操作
func LoginByApp(userID string, tokenKey string) (string, error) {
	token, err := authentication.genToken(userID, tokenKey)
	if err != nil {
		return "", err
	} else {
		return token, nil
	}
}

// Logout 注销操作
func Logout(r *gin.Context) {
	r.SetCookie(tokenName, "", -1, "/", "", false, false)
	token := authentication.getTokenByRequest(r)
	if len(token) > 0 {
		Del(token)
	}
}
