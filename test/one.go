package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	router := gin.Default()
	router.GET("/ping", ping) //注意这里的ping是一个函数对象，不是调用ping函数
	err := router.Run()
	if err != nil {
		return
	}
}
