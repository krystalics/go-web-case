package api

import (
	"github.com/gin-gonic/gin"
	"go-web-case/internal/app/go-web-case/common"
	"go-web-case/internal/app/go-web-case/model"
	"go-web-case/internal/app/go-web-case/service"
)

func CreateUser(c *gin.Context) {
	var user model.User
	c.Bind(&user)

	service.Srv.CreateUser(user)
	common.ResSuccess(c, nil)
}
