package main

import "fmt"

func main() {
	name := "dj"
	age := 22
	info := fmt.Sprintf("姓名:%s,年龄:%d", name, age)
	fmt.Println(info)
	fmt.Println(fmt.Sprintf("姓名:%s,年龄:%d", "DongJun", 23))
}
