package web

import (
	"embed"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/remoting/frame/server/auth"
)

type HandlerFunc func(*Context)
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

func (engine *Engine) Use(middleware ...gin.HandlerFunc) {
	engine.Engine.Use(middleware...)
}
func (engine *Engine) SetMode(mode string) {
	gin.SetMode(mode)
}
func (engine *Engine) Static(prefix string, fs embed.FS) {
	engine.Any(prefix+"/*filepath", func(c *Context) {
		staticServer := http.FileServer(http.FS(fs))
		staticServer.ServeHTTP(c.Writer, c.Request)
	})
	engine.GET("/", func(c *Context) {
		c.Redirect(301, prefix+"/")
	})
}
func (engine *Engine) Any(relativePath string, handlerFunc HandlerFunc) {
	engine.Engine.Any(relativePath, func(c *gin.Context) {
		handlerFunc(&Context{
			Context: c,
		})
	})
}
func (engine *Engine) GET(relativePath string, handlerFunc HandlerFunc) {
	engine.Engine.GET(relativePath, func(c *gin.Context) {
		handlerFunc(&Context{
			Context: c,
		})
	})
}
func (engine *Engine) POST(relativePath string, handlerFunc HandlerFunc) {
	engine.Engine.POST(relativePath, func(c *gin.Context) {
		handlerFunc(&Context{
			Context: c,
		})
	})
}
func (engine *Engine) Api(prefix string, controllers map[string]any) *RouterGroup {
	// API 路由
	api := engine.Group(prefix)
	api.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"time": time.Now().Unix(),
		})
	})
	apix := &RouterGroup{
		RouterGroup: api,
	}
	for prefix, c := range controllers {
		apix.AddRouter(prefix, c)
	}
	return apix
}
func (engine *Engine) Add(prefix string, controllers map[string]Controller) *RouterGroup {
	// API 路由
	api := engine.Group(prefix)
	api.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"time": time.Now().Unix(),
		})
	})
	apix := &RouterGroup{
		RouterGroup: api,
	}
	for prefix, c := range controllers {
		apix.AddRouter(prefix, c)
	}
	return apix
}
