package router

import (
	"github.com/gin-gonic/gin"
	"go-web-case/internal/app/go-web-case/api"
)

func RegisterUser(e *gin.Engine) {
	routers := []URLHandlerMap{
		// hello
		{"POST", "/api/user/create", api.CreateUser},
	}

	RouterFill(&e.RouterGroup, &routers)
}
