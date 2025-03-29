package conf

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

func ReadConfigFile(filename string) *viper.Viper {
	config := viper.New()
	//从项目的根目录下读取名为config，格式为yaml的文件
	config.SetConfigName(filename)
	config.SetConfigType("yaml")
	config.AddConfigPath(".")
	err := config.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			panic(fmt.Errorf("配置文件不存在：%w", err))
		} else {
			panic(fmt.Errorf("读取配置文件错误：%w", err))
		}
	}
	return config
}
