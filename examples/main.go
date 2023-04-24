package main

import (
	"net/http"

	"github.com/remoting/frame/examples/controller"
	"github.com/remoting/frame/examples/service"
	"github.com/remoting/frame/server/auth"
	"github.com/remoting/frame/server/web"
	"github.com/remoting/frame/spring"
)

func main() {
	// 创建框架
	app := web.New()
	//app.SetMode("release")
	// 认证体系
	app.Use(auth.GetAuthHandlerFunc(authConfig()))
	// 开启静态文件服务
	app.StaticFS("/ui", http.Dir("./ui"))
	// 创建API路由
	app.Api("/api", controller.GetControllers())
	app.Run(":8080")
}

func authConfig() auth.AuthConfig {
	return auth.AuthConfig{
		TokenSecret: "text",
		AnonymousPath: []string{
			"/api/user/Login",
			"/api/user/Logout",
			"/api/test/Login",
			"/ui",
			"/api/info",
			"/favicon.ico",
		},
		UserService: spring.GetBean[*service.UserService](controller.GetSpring()),
	}
}
