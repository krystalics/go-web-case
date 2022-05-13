package handlerServer

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"net/http/httputil"
	"time"
)

func Ping(c *gin.Context) {

	go func() {
		zap.L().Info("do something async start")
		time.Sleep(time.Second * 1)
		zap.L().Info("do something async done")
	}()

	c.String(http.StatusOK, "pong")

}

func Ping2(c *gin.Context) {

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
}
