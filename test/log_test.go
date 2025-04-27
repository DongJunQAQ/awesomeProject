package main

import (
	"awesomeProject/conf"
	"fmt"
	"testing"
)

func TestGetGlobalLogger(t *testing.T) { //测试全局日志实例是否为单例模式
	logger1 := conf.GetGlobalLogger()
	fmt.Printf("%p\n", logger1)
	logger2 := conf.GetGlobalLogger()
	fmt.Printf("%p\n", logger2)
	if logger1 != logger2 {
		t.Fail()
		fmt.Println("并非单例模式")
	} else {
		fmt.Println("是单例模式")
	}
}
