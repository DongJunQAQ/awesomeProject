package main

import "fmt"

func lastElement[T any](slice []T) T { // 定义一个泛型函数，用于获取切片的最后一个元素
	return slice[len(slice)-1]
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5} // 定义一个int类型的切片
	intLast := lastElement(intSlice) // 调用泛型函数获取int切片的最后一个元素
	fmt.Println("int切片的最后一个元素是：", intLast)
	stringSlice := []string{"apple", "banana", "cherry"} // 定义一个string类型的切片
	stringLast := lastElement(stringSlice)               // 调用泛型函数获取string切片的最后一个元素
	fmt.Println("string切片的最后一个元素是：", stringLast)
}
