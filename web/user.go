package web

import (
	"awesomeProject/conf"
	"awesomeProject/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

var db = database.GetGlobalDB()
var logger = conf.GetGlobalLogger()

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func QueryAllUsers(c *gin.Context) { //指定字段查询所有用户
	users := database.SelectFieldsQueryAllUser(db)
	c.JSON(http.StatusOK, users)
	logger.Debugf("URL:%s Code:%d", c.Request.URL.String(), c.Writer.Status())
}
