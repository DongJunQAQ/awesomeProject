package conf

import (
	"go.uber.org/zap"
)

func InitDevLogger() {
	logger, _ := zap.NewDevelopment()
	logger.Info("dev this is info",
		zap.String("name", "DongJun"),
		zap.Uint8("age", 22),
		zap.Bool("isDelete", false))
}
