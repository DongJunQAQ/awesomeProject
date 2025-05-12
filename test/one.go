package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

func Hello(c *gin.Context) { // 处理函数1
	c.String(http.StatusOK, "Hello, World!")
}

func Goodbye(c *gin.Context) { // 处理函数2
	c.String(http.StatusOK, "Goodbye!")
}

func ReflectHandler(c *gin.Context) { // 反射处理函数
	path := c.Param("path")                              // 获取请求路径参数
	funcValue := reflect.ValueOf(map[string]interface{}{ // 构建一个键为string类型，值为interface{}任意类型的map，然后reflect.ValueOf函数会将funcValue变量表示为某个函数，其存储的是该函数的反射表示，后续就能通过反射机制调用这个函数
		"hello":   Hello, //值的类型为函数
		"goodbye": Goodbye,
	}[path]) //借助前面获取的path变量作为键，从map中取出对应的函数值。如若path为"hello"，则取出Hello函数
	//funcValue是一个reflect.Value类型的对象，它是函数的反射表示其包含了该函数的相关信息，比如函数的类型、可以调用该函数等，其并非直接存储函数的地址，而是对函数的一种抽象表示，提供了操作该函数的接口
	if funcValue.IsValid() == false { // 如果该函数的反射表示无效的话
		c.AbortWithStatus(http.StatusNotFound) //则当前请求的处理流程会立即终止，并返回404且后续注册的中间件或处理函数将不会被执行
		return
	}
	args := []reflect.Value{ // 准备参数
		reflect.ValueOf(c),
	}
	funcValue.Call(args) // 调用函数并传入参数c
}

func main() {
	r := gin.Default()
	r.GET("/:path", ReflectHandler) // 注册反射处理函数，这里的:path属于动态路由参数，当有请求发送到如http://127.0.0.1:8080/goodbye时，path变量就会被赋值为 "goodbye"
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
