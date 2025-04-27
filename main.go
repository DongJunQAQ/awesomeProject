package main

import (
	"awesomeProject/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var myConf = conf.GetGlobalConf() //获取配置

func connectDB() *gorm.DB { //连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", myConf.GetString("db.user"), myConf.GetString("db.passwd"), myConf.GetString("db.host"), myConf.GetString("db.port"), myConf.GetString("db.database"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("无法连接到数据库")
	}
	return db
}

func main() {
	//db := connectDB()
	//err := db.AutoMigrate(&model.User{}) //自动迁移数据库
	//if err != nil {
	//	panic("数据库迁移失败")
	//}
	//fmt.Println(database.CreateManyRecord(db))

	logger := conf.GetGlobalLogger()
	logger.Debugf("报错级别:%s", "Debug")
	logger.Info("qwe")
	logger.Warn("qwe")
	logger.Error("qwe")
}
