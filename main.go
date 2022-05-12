package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-web-case/config"
	"go-web-case/internal/app/go-web-case/conf"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	InitApp()
	gin.SetMode(viper.GetString("gin.mode"))
	router := gin.Default()
	//全局使用跨域的配置、改造了gin的logger 以及对应的recovery
	router.Use(conf.CorsHandler(), conf.GinLogger(), conf.GinRecovery(true))

	//中间需要将 路径和对应的处理方法绑定
	router.GET("/ping", func(c *gin.Context) {
		go c.String(http.StatusOK, "pong")
	})

	router.Run("127.0.0.1:8080") // listen and serve on 0.0.0.0:8080
}

func InitApp() {
	config.InitConf()
	err := conf.InitLogger(config.GetLogConfig())
	if err != nil {
		zap.L().Error("init logger error")
		return
	}
	err = conf.InitDB(config.GetDataSourceConfig())
	if err != nil {
		zap.L().Error("init db error")
		return
	}
}
