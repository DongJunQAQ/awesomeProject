package conf

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var (
	globalConf *viper.Viper //使用单例模式实现配置实例
	confOnce   sync.Once
)

func readConfigFile(filename string) *viper.Viper {
	confOnce.Do(func() {
		globalConf = viper.New()
		//从项目的根目录下读取名为config，格式为yaml的文件
		globalConf.SetConfigName(filename)
		globalConf.SetConfigType("yaml")
		globalConf.AddConfigPath(".")
		err := globalConf.ReadInConfig()
		if err != nil {
			var configFileNotFoundError viper.ConfigFileNotFoundError
			if errors.As(err, &configFileNotFoundError) {
				panic(fmt.Errorf("配置文件%s不存在:%w", filename, err))
			} else {
				panic(fmt.Errorf("读取配置文件%s错误：%w", filename, err))
			}
		}
	})
	return globalConf
}
func GetGlobalConf() *viper.Viper { //获取全局单例配置实例
	return readConfigFile("config") //指定配置文件的文件名
}
