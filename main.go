package main

import (
	"awesomeProject/conf"
	"awesomeProject/database"
	"awesomeProject/model"
	"awesomeProject/web"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化数据库
	db := database.GetGlobalDB()
	err := db.AutoMigrate(&model.User{}) //自动迁移数据库
	if err != nil {
		conf.GetGlobalLogger().Panicf("数据库迁移失败")
	}
	//初始化web
	router := gin.Default()       //创建默认的Gin路由实例
	router.GET("/ping", web.Ping) //这里web.Ping与web.Ping()的区别
	err = router.Run()            //启动web服务且默认监听8080端口
	if err != nil {
		conf.GetGlobalLogger().Panicf("web服务启动失败:%s", err)
	}
}
