package router

import (
	"github.com/gin-gonic/gin"
	"go-web-case/internal/app/go-web-case/handler"
)

func RegisterHello(e *gin.Engine) {
	routers := []URLHandlerMap{
		// hello
		{"GET", "/ping", handler.Ping},
		{"GET", "/ping2", handler.Ping2},
	}

	RouterFill(&e.RouterGroup, &routers)
}
