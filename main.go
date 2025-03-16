package main

import (
	"awesomeProject/database"
	"awesomeProject/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectDB() *gorm.DB {
	dsn := "root:123456@tcp(192.168.246.137:3306)/goweb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("无法连接到数据库")
	}
	return db
}

func main() {
	db := connectDB()
	err := db.AutoMigrate(&model.User{}) //自动迁移
	if err != nil {
		panic("数据库迁移失败")
	}
	fmt.Printf("已成功创建%d条记录\n", database.CreateOneRecord(db))
	fmt.Printf("已成功创建%d条记录\n", database.CreateManyRecord(db))
	fmt.Printf("已成功创建%d条记录\n", database.SelectedFields(db))
	//fmt.Printf("已成功创建%d条记录\n", database.IgnoreFields(db))
}
