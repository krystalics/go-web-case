package router

import (
	"github.com/gin-gonic/gin"
	"go-web-case/internal/app/go-web-case/controller"
)

func RegisterHello(e *gin.Engine) {
	routers := []URLHandlerMap{
		// hello
		{"GET", "/ping", controller.Ping},
		{"GET", "/ping2", controller.Ping2},
	}

	RouterFill(&e.RouterGroup, &routers)
}
