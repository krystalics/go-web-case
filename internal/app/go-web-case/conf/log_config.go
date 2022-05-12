package conf

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http"
	"net/http/httputil"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"max_size"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

var globalLogger *zap.Logger

//todo 全局对象、需要程序启动前调用这个Init方法
//需要先从配置文件获取这些配置
func InitLogger(cfg *LogConfig) (err error) {
	writer := getLogWriter(cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	encoder := getEncoder()

	l := new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return err
	}

	core := zapcore.NewCore(
		encoder, writer, l,
	)

	globalLogger = zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel), //error级别的日志打印堆栈
	)

	zap.ReplaceGlobals(globalLogger) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可

	return
}

//lumberJack 可以对日志进行滚动写入
func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,   //最大10M
		MaxAge:     maxAge,    //老文件最多30天
		MaxBackups: maxBackup, //最多老文件数
		LocalTime:  false,
		Compress:   false,
	}

	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(config)

}

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
		globalLogger.Info(path,
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

//gin会有自己的recover、但是日志是打印在console、需要和zap配合的话也要重写
//recover可能出现的panic
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
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
					globalLogger.Error(
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
					globalLogger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(request)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					globalLogger.Error("[Recovery from panic]",
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
