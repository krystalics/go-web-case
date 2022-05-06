package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-web-case/configs"
	"go-web-case/internal/app/go-web-case/config"
)

type App struct {
}

func main() {
	configs.InitConf()
	config.InitLogger(&config.LogConfig{
		Level:      viper.GetString("log.level"),
		Filename:   viper.GetString("log.filename"),
		MaxSize:    viper.GetInt("log.maxSize"),
		MaxAge:     viper.GetInt("log.maxAge"),
		MaxBackups: viper.GetInt("log.maxBackups"),
	})

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	//全局使用跨域的配置、改造了gin的logger
	router.Use(config.CorsHandler(), config.GinLogger(), config.GinRecovery(true))

	//中间需要将 路径和对应的处理方法绑定
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run("127.0.0.1:8080") // listen and serve on 0.0.0.0:8080
}
