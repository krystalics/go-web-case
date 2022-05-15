package router

import (
	"github.com/gin-gonic/gin"
	"go-web-case/internal/app/go-web-case/middleware"
	"net/http"
)

type URLHandlerMap struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	var router = gin.Default()
	//添加业务路由，根据需要选择需要的中间件
	router.Use(middleware.CorsHandler(), middleware.GinRecovery(true), middleware.GinLogger())

	RegisterPprof(router)
	RegisterHello(router)

	return router
}

// GinHandler 将http.HandlerFunc转为gin.HandlerFunc
func GinHandler(h http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func RouterFill(router *gin.RouterGroup, routers *[]URLHandlerMap) {
	for _, r := range *routers {
		router.Handle(r.Method, r.Path, r.Handler)
	}
}
