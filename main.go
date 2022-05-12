package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-web-case/internal/app/go-web-case/conf"
	"go.uber.org/zap"
	"net/http"
	"net/http/httputil"
	"time"
)

func main() {
	InitApp()
	gin.SetMode(viper.GetString("gin.mode"))
	router := gin.Default()
	//全局使用跨域的配置、改造了gin的logger 以及对应的recovery
	router.Use(conf.CorsHandler(), conf.GinRecovery(true), conf.GinLogger())

	//中间需要将 路径和对应的处理方法绑定
	router.GET("/ping", func(c *gin.Context) {

		go func() {
			zap.L().Info("do something async start")
			time.Sleep(time.Second * 1)
			zap.L().Info("do something async done")
		}()

		c.String(http.StatusOK, "pong")
	})

	router.GET("/ping2", func(c *gin.Context) {

	})

	router.GET("/ping3", func(c *gin.Context) {

		//捕获到这个panic
		go func() {
			defer func() {
				request, _ := httputil.DumpRequest(c.Request, false)
				e := recover()
				if e != nil {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", e),
						zap.String("request", string(request)),
					)
				}
			}()

			c.String(http.StatusOK, "pong")
		}()
	})

	zap.L().Info("app start at 127.0.0.1:8080")
	err := router.Run("127.0.0.1:8080")
	if err != nil {
		zap.L().Fatal("app start failed ", zap.Error(err))
		return
	}

}

func InitApp() {
	conf.InitConf()
	err := conf.InitLogger(conf.GetLogConfig())
	if err != nil {
		zap.L().Error("init logger error")
		return
	}
	err = conf.InitDB(conf.GetDataSourceConfig())
	if err != nil {
		zap.L().Error("init db error")
		return
	}
	zap.L().Info("app config init success")
}
