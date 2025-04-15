package main

import "fmt"

func PrintValue(v interface{}) { //接收一个空接口（任意类型）类型的参数
	fmt.Printf("值为:%v，类型为:%T\n", v, v)
}

func main() {
	// 定义不同类型的变量
	num := 10
	str := "hello"
	arr := [3]int{1, 2, 3}
	// 调用PrintValue函数并传入不同类型的变量
	PrintValue(num)
	PrintValue(str)
	PrintValue(arr)
}
