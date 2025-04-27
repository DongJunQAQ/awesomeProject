package main

import (
	"awesomeProject/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.MyConf.GetString("db.user"), conf.MyConf.GetString("db.passwd"), conf.MyConf.GetString("db.host"), conf.MyConf.GetString("db.port"), conf.MyConf.GetString("db.database"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("无法连接到数据库")
	}
	return db
}

func main() {
	//db := connectDB()
	//err := db.AutoMigrate(&model.User{}) //自动迁移
	//if err != nil {
	//	panic("数据库迁移失败")
	//}
	//fmt.Println(database.CreateManyRecord(db))
	logger := conf.GetProdLogger()
	logger.Debug("qwe")
	logger.Info("qwe")
	logger.Warn("qwe")
	logger.Error("qwe")
}
