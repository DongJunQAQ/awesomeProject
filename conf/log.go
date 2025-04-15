package conf

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const ( //定义颜色转义码
	Blue   = "\033[34m"
	Yellow = "\033[33m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	ReSet  = "\033[0m" //每行日志输出后需要还原至默认颜色
)

func encodeLevel(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
	switch level {
	case zapcore.DebugLevel:
		encoder.AppendString(Green + "DEBUG" + ReSet)
	case zapcore.InfoLevel:
		encoder.AppendString(Blue + "INFO" + ReSet)
	case zapcore.WarnLevel:
		encoder.AppendString(Yellow + "WARNING" + ReSet)
	case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		encoder.AppendString(Red + "ERROR" + ReSet)
	default:
		encoder.AppendString(level.String())
	}
}

func InitDevLogger() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = encodeLevel
	logger, _ := cfg.Build()
	logger.Debug("dev this is debug")
	logger.Info("dev this is info")
	logger.Warn("dev this is warn")
	logger.Error("dev this is error")
}
