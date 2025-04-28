package database

import (
	"awesomeProject/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	myConf   = conf.GetGlobalConf() //获取全局配置实例
	dbOnce   sync.Once
	GlobalDB *gorm.DB //全局数据库实例
)

func GetGlobalDB() *gorm.DB { //获取全局数据库单例实例
	dbOnce.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", myConf.GetString("db.user"), myConf.GetString("db.passwd"), myConf.GetString("db.host"), myConf.GetString("db.port"), myConf.GetString("db.database"))
		var err error
		GlobalDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			conf.GetGlobalLogger().Panicf("无法连接至数据库:%s:%s", myConf.GetString("db.host"), myConf.GetString("db.port"))
		} else {
			conf.GetGlobalLogger().Debugf("数据库连接成功")
		}
	})
	return GlobalDB
}
