package main

import (
	"awesomeProject/database"
	"fmt"
	"testing"
)

var db = database.GetGlobalDB()

func TestGetGlobalDB(t *testing.T) {
	db1 := database.GetGlobalDB()
	fmt.Printf("%p\n", db1)
	db2 := database.GetGlobalDB()
	fmt.Printf("%p\n", db2)
	if db1 != db2 {
		t.Fail()
		fmt.Println("并非单例模式")
	} else {
		fmt.Println("是单例模式")
	}
}

func TestCreateOneRecord(t *testing.T) {
	fmt.Println(database.CreateOneRecord(db))
}

func TestSelectFieldsQueryAllUser(t *testing.T) {
	fmt.Println(database.SelectFieldsQueryAllUser(db))
}
