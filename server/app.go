package server

import (
	"embed"
	"fmt"
	"github.com/remoting/frame/pkg/logger"
	"github.com/remoting/frame/server/tools"
	"github.com/remoting/frame/server/web"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Engine struct {
	*gin.Engine
}

func New() *Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		fmt.Printf("[%s] [DEBUG] %-6s %-25s --> %s (%d handlers)\n", logger.Conf.Prefix, httpMethod, absolutePath, handlerName, nuHandlers)
	}
	r := gin.New()
	gin.SetMode(gin.DebugMode)
	r.Use(tools.Logger())
	r.Use(tools.ErrorHandler())
	r.Use(tools.CORSMiddleware())
	//r.Use(auth.Auth())

	return &Engine{
		Engine: r,
	}
}
func (engine *Engine) Run(addr ...string) (err error) {
	gin.SetMode(gin.ReleaseMode)
	return engine.Engine.Run(addr...)
}
func (engine *Engine) Use(middleware ...gin.HandlerFunc) {
	engine.Engine.Use(middleware...)
}
func (engine *Engine) SetMode(mode string) {
	gin.SetMode(mode)
}
func (engine *Engine) Static(prefix string, fs embed.FS) {
	engine.Any(prefix+"/*filepath", func(c *web.Context) {
		staticServer := http.FileServer(http.FS(fs))
		staticServer.ServeHTTP(c.Writer, c.Request)
	})
}
func (engine *Engine) Any(relativePath string, handlerFunc web.HandlerFunc) {
	engine.Engine.Any(relativePath, func(c *gin.Context) {
		handlerFunc(&web.Context{
			Context: c,
		})
	})
}
func (engine *Engine) GET(relativePath string, handlerFunc web.HandlerFunc) {
	engine.Engine.GET(relativePath, func(c *gin.Context) {
		handlerFunc(&web.Context{
			Context: c,
		})
	})
}
func (engine *Engine) POST(relativePath string, handlerFunc web.HandlerFunc) {
	engine.Engine.POST(relativePath, func(c *gin.Context) {
		handlerFunc(&web.Context{
			Context: c,
		})
	})
}
func (engine *Engine) Api(prefix string, controllers map[string]any) *web.RouterGroup {
	// API 路由
	api := engine.Group(prefix)
	api.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"time": time.Now().Unix(),
		})
	})
	apix := &web.RouterGroup{
		RouterGroup: api,
	}
	for prefix, c := range controllers {
		apix.AddRouter(prefix, c)
	}
	return apix
}

func (engine *Engine) Add(prefix string, controllers ...web.Controller) *web.RouterGroup {
	//// API 路由
	api := engine.Group(prefix)
	apix := &web.RouterGroup{
		RouterGroup: api,
	}
	for _, controller := range controllers {
		controller.OnInit(func(_prefix string) *web.RouterGroup {
			return &web.RouterGroup{
				RouterGroup: apix.Group(_prefix),
			}
		})
	}
	return apix
}
