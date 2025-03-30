package conf

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitDevLogger() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") //设置该配置的日志时间格式
	logger, _ := cfg.Build()
	logger.Debug("dev this is debug")
	logger.Info("dev this is info")
	logger.Warn("dev this is warn")
	logger.Error("dev this is error")
}
