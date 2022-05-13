package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

//gin有自己的日志体系、和zap需要融合
//参考 https://www.liwenzhou.com/posts/Go/use_zap_in_gin/
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		//执行后续中间件
		c.Next()
		//执行完成后，属于是整个middleware set都执行完后、才会按照调用栈都顺序执行后面的代码

		cost := time.Since(start)
		zap.L().Info(path,
			zap.String("routine", strconv.Itoa(runtime.NumGoroutine())),
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

//GinRecovery gin会有自己的recover、但是日志是打印在console、需要和zap配合的话也要重写
//recover可能出现的panic
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		//panic: http: wrote more than the declared Content-Length 出现panic没有捕捉到
		//在middleware最里层就是handler了，所以这个理论上能捕获handler的panic

		//利用recover处理panic指令，defer必须在panic之前声明，否则当panic时，recover无法捕获到panic．
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if se, ok := err.(*net.OpError); ok {
					if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
						brokenPipe = true
					}
				}
				request, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(
						c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(request)),
					)
					//if the connection is dead,we can't write a status to it
					c.Error(err.(error))
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(request)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(request)),
					)
				}

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		c.Next()
	}
}
