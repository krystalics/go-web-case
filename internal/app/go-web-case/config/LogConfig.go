package config

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.SugaredLogger

//todo 全局对象、需要程序启动前调用这个Init方法
func InitLogger() {
	writer := getLogWriter()
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)

	log := zap.New(core, zap.AddCaller())

	globalLogger = log.Sugar()
}

//lumberJack 可以对日志进行滚动写入
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    1,
		MaxAge:     30,
		MaxBackups: 5,
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
