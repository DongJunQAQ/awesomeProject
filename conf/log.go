package conf

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func encodeLevel(zapcore.Level, zapcore.PrimitiveArrayEncoder) {

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
