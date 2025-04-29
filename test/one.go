package main

import (
	"awesomeProject/conf"
	"github.com/gin-gonic/gin"
	"net/http"
)

var myLogger = conf.GetGlobalLogger()

func ping(c *gin.Context) { //获取响应体中的相关信息(请求方法、请求URL、响应状态码)
	c.JSON(http.StatusOK, gin.H{"msg": "get pong"})
	myLogger.Debugf("Method: %s, Path: %s, Status: %d", c.Request.Method, c.Request.URL.String(), c.Writer.Status())
}

func main() {
	router := gin.Default()
	router.GET("/ping", ping)
	err := router.Run()
	if err != nil {
		return
	}
}
