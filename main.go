package main

import (
	"github.com/gin-gonic/gin"
	"go-web-case/internal/app/go-web-case/conf"
	handlerServer "go-web-case/internal/app/go-web-case/handler"
	"go-web-case/internal/app/go-web-case/middleware"
	ginpprof "go-web-case/internal/app/go-web-case/pprof"
	"go.uber.org/zap"
	"strconv"
)

func main() {
	config := conf.New()
	conf.InitLogger(conf.GetLogConfig())
	zap.L().Info("app config init success")

	gin.SetMode(config.App.RunMode)
	router := gin.Default()
	router.Use(middleware.CorsHandler(), middleware.GinRecovery(true), middleware.GinLogger())

	handlerServer.New(config, router)
	ginpprof.Wrap(router)

	addr := "127.0.0.1:" + strconv.Itoa(config.App.HttpPort)
	zap.L().Info("app start at " + addr)
	err := router.Run(addr)
	if err != nil {
		zap.L().Fatal("app start failed ", zap.Error(err))
		return
	}

}
