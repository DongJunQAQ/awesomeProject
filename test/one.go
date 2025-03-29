package main

import (
	"awesomeProject/conf"
	"fmt"
)

func main() {
	myConf := conf.ReadConfigFile("config")
	fmt.Println(myConf.GetString("db.host"))
	fmt.Println(myConf.GetString("db.port"))
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", myConf.GetString("db.user"), myConf.GetString("db.passwd"), myConf.GetString("db.host"), myConf.GetString("db.port"), myConf.GetString("db.database"))
	fmt.Println(dsn)
}
