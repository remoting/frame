package web

import (
	"embed"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/remoting/frame/server/auth"
)

type Engine struct {
	*gin.Engine
}

func New() *Engine {
	r := gin.Default()
	r.Use(auth.ErrorHandler())
	r.Use(auth.CORSMiddleware())
	//r.Use(auth.Auth())

	return &Engine{
		Engine: r,
	}
}
func (engine *Engine) SetMode(mode string) {
	gin.SetMode(mode)
}
func (engine *Engine) Static(prefix string, fs embed.FS) {
	engine.Any(prefix+"/*filepath", func(c *gin.Context) {
		staticServer := http.FileServer(http.FS(fs))
		staticServer.ServeHTTP(c.Writer, c.Request)
	})
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(301, prefix+"/")
	})
}
func (engine *Engine) Api(prefix string) *RouterGroup {
	// API 路由
	api := engine.Group(prefix)
	api.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"time": time.Now().Unix(),
		})
	})
	return &RouterGroup{
		RouterGroup: api,
	}
}
