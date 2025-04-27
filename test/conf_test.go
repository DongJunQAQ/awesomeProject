package main

import (
	"awesomeProject/conf"
	"fmt"
	"testing"
)

func TestGetGlobalConf(t *testing.T) { //测试全局配置实例是否为单例模式
	conf1 := conf.GetGlobalConf()
	fmt.Printf("%p\n", conf1)
	conf2 := conf.GetGlobalConf()
	fmt.Printf("%p\n", conf2)
	if conf1 != conf2 {
		t.Fail()
		fmt.Println("并非单例模式")
	} else {
		fmt.Println("是单例模式")
	}
}
