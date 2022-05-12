package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-web-case/config"
	"go-web-case/internal/app/go-web-case/conf"
	"go.uber.org/zap"
	"net/http"
)

type App struct {
}

func main() {
	config.InitConf()
	err := conf.InitLogger(&conf.LogConfig{
		Level:      viper.GetString("log.level"),
		Filename:   viper.GetString("log.filename"),
		MaxSize:    viper.GetInt("log.maxSize"),
		MaxAge:     viper.GetInt("log.maxAge"),
		MaxBackups: viper.GetInt("log.maxBackups"),
	})

	if err != nil {
		zap.L().Error("init error")
		return
	}

	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	//全局使用跨域的配置、改造了gin的logger 以及对应的recovery
	router.Use(conf.CorsHandler(), conf.GinLogger(), conf.GinRecovery(true))

	//中间需要将 路径和对应的处理方法绑定
	router.GET("/ping", func(c *gin.Context) {
		go c.String(http.StatusOK, "pong")
	})

	router.Run("127.0.0.1:8080") // listen and serve on 0.0.0.0:8080
}
