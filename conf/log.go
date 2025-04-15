package conf

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

var GlobalLogger *zap.Logger //单例模式
var Once sync.Once

func GetProdLogger() *zap.Logger {
	Once.Do(func() {
		cfg := zap.NewProductionConfig()
		consoleCore := zapcore.NewCore(zapcore.NewJSONEncoder(cfg.EncoderConfig), os.Stdout, zapcore.DebugLevel) //控制台输出日志
		file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		fileCore := zapcore.NewCore(zapcore.NewJSONEncoder(cfg.EncoderConfig), file, zapcore.DebugLevel) //输出日志文件并设置Debug级别为低显示级别
		core := zapcore.NewTee(consoleCore, fileCore)
		GlobalLogger = zap.New(core, zap.AddCaller())
	})
	return GlobalLogger
}
