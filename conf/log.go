package conf

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

var (
	sugarGlobalLogger *zap.SugaredLogger //使用单例模式实现日志实例
	loggerOnce        sync.Once
)

func convertLevelFormat(confLevel string) zapcore.Level { //将配置文件中的日志级别转换为zapcore.Level
	switch confLevel {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func GetGlobalLogger() *zap.SugaredLogger {
	loggerOnce.Do(func() {
		var GlobalLogger *zap.Logger
		cfg := zap.NewProductionConfig()
		zapLogLevel := convertLevelFormat(GetGlobalConf().GetString("log.level")) //获取zap格式的日志级别
		cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
		consoleCore := zapcore.NewCore(zapcore.NewJSONEncoder(cfg.EncoderConfig), os.Stdout, zapLogLevel) //控制台输出日志
		file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		fileCore := zapcore.NewCore(zapcore.NewJSONEncoder(cfg.EncoderConfig), file, zapLogLevel) //输出日志文件
		core := zapcore.NewTee(consoleCore, fileCore)
		GlobalLogger = zap.New(core, zap.AddCaller())
		sugarGlobalLogger = GlobalLogger.Sugar()
	})
	return sugarGlobalLogger
}
