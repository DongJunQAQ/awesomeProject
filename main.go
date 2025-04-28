package main

import (
	"awesomeProject/conf"
	"awesomeProject/database"
	"awesomeProject/model"
)

func main() {
	db := database.GetGlobalDB()
	err := db.AutoMigrate(&model.User{}) //自动迁移数据库
	if err != nil {
		conf.GetGlobalLogger().Panicf("数据库迁移失败")
	}
	//fmt.Println(database.SelectFieldsQuery(db))
}
