package conf

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"max_size"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

func InitLogger(cfg *Config) {
	writer := getLogWriter(cfg.Log.Filename, cfg.Log.MaxSize, cfg.Log.MaxBackups, cfg.Log.MaxAge)
	encoder := getEncoder()

	l := new(zapcore.Level)
	err := l.UnmarshalText([]byte(cfg.Log.Level))
	if err != nil {
		return
	}

	core := zapcore.NewCore(
		encoder, writer, l,
	)

	var logger *zap.Logger
	if cfg.App.RunMode == "debug" {
		logger, _ = zap.NewDevelopment()
	} else {
		logger = zap.New(
			core,
			zap.AddCaller(),
			zap.AddStacktrace(zap.ErrorLevel), //error级别的日志打印堆栈
		)
	}

	zap.ReplaceGlobals(logger) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
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
