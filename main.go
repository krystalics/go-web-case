package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-web-case/internal/app/go-web-case/conf"
	handlerServer "go-web-case/internal/app/go-web-case/handler"
	"go-web-case/internal/app/go-web-case/router"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	config := conf.New()
	conf.InitLogger(conf.GetLogConfig())
	gin.SetMode(config.App.RunMode)
	zap.L().Info("app config init success")
	r := router.InitRouter()
	handlerServer.New(config, r)

	serverAddr := "127.0.0.1:" + strconv.Itoa(config.App.HttpPort)
	//TODO ReadTimeout 和 WriteTimeout根据实际需要去修改
	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      r,
		ReadTimeout:  time.Duration(config.App.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.App.WriteTimeout) * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		zap.L().Fatal("listen error", zap.Error(err))
	}
	zap.L().Info("Server start", zap.String("addr", "127.0.0.1"), zap.String("port", serverAddr))

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Error("Shutting down server...")

	// The context is used to inform the server it has 10 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server forced to shutdown:", zap.Error(err))
	}
	zap.L().Info("Server exiting")

}
