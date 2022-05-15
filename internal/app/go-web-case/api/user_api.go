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

	_, err := service.Srv.CreateUser(user)
	if err != nil {
		common.ResFailed(c, err)
	} else {
		common.ResSuccess(c, nil)
	}
}
