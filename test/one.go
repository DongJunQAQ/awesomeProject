package main

import "fmt"

func myFunc(value interface{}) { //接收一个任意类型的实参
	fmt.Printf("数据%v的类型为:%T\n", value, value)
}

func myFunc2(value ...interface{}) { //接收多个任意类型的实参
	fmt.Printf("数据%v的类型为:%T\n", value, value)
	for _, v := range value {
		myFunc(v)
	}
}

func main() {
	myFunc(2)
	myFunc("qwe")
	fmt.Println()
	myFunc2(1, "qwe")
}
