package handler

import (
	"github.com/gin-gonic/gin"
	"go-web-case/internal/app/go-web-case/conf"
	"go-web-case/internal/app/go-web-case/service"
)

//类似于替代spring中的controller层
var srv *service.Service

func New(c *conf.Config, e *gin.Engine) {
	srv = service.New(c)
}
