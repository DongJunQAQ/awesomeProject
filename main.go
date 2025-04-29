package main

import (
	"awesomeProject/conf"
	"awesomeProject/database"
	"awesomeProject/model"
	"awesomeProject/web"
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	myConf   = conf.GetGlobalConf()
	myLogger = conf.GetGlobalLogger()
)

func main() {
	//初始化数据库
	db := database.GetGlobalDB()
	err := db.AutoMigrate(&model.User{}) //自动迁移数据库
	if err != nil {
		myLogger.Panicln("数据库迁移失败:", err)
	} else {
		myLogger.Infoln("数据库迁移成功")
	}
	//初始化web
	router := gin.Default()           //创建默认的Gin路由实例
	router.Use(web.LoggingMiddleware) //全局中间件，每个handler执行之前都会执行该中间件
	router.GET("/ping", web.Ping)     //web.Ping代表的是一个函数对象
	router.GET("/user", web.QueryAllUsers)
	err = router.Run(fmt.Sprintf("%s:%s", myConf.GetString("web.host"), myConf.GetString("web.port"))) //启动web服务且监听端口
	if err != nil {
		conf.GetGlobalLogger().Panicf("web服务启动失败:%s", err)
	}
}
