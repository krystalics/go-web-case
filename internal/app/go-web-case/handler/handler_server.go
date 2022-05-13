package handlerServer

import (
	"github.com/gin-gonic/gin"
	"go-web-case/internal/app/go-web-case/conf"
	"go-web-case/internal/app/go-web-case/service"
)

//类似于替代spring中的controller层

type URLHandlerMap struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

var srv *service.Service

func New(c *conf.Config, e *gin.Engine) {
	srv = service.New(c)
	routers := buildRouters()
	route(&e.RouterGroup, routers)
}

func route(router *gin.RouterGroup, routers *[]URLHandlerMap) {
	for _, r := range *routers {
		router.Handle(r.Method, r.Path, r.Handler)
	}
}

func buildRouters() *[]URLHandlerMap {
	return &[]URLHandlerMap{
		// hello
		{"GET", "/ping", Ping},
		{"GET", "/ping2", Ping2},
	}

}
