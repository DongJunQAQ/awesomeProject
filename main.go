package main

import (
	"awesomeProject/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var myConf = conf.ReadConfigFile("config")

func connectDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", myConf.GetString("db.user"), myConf.GetString("db.passwd"), myConf.GetString("db.host"), myConf.GetString("db.port"), myConf.GetString("db.database"))
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
	//fmt.Println(database.SelectFieldsQuery(db))
	myProdLogger := conf.GetProdLogger()
	myProdLogger.Debug("prod this is Debug")
	myProdLogger.Info("prod this is Info")
	myProdLogger.Warn("prod this is Warn")
	myProdLogger.Error("prod this is Error")
}
