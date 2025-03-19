package database

import (
	"awesomeProject/model"
	"fmt"
	"gorm.io/gorm"
)

type ApiUser struct { //创建指定要返回字段的结构体
	Name string
	Age  uint8
}

func SelectFieldsQuery(db *gorm.DB) []ApiUser { //选择指定字段检索返回
	var users []ApiUser
	result := db.Model(&model.User{}).Select("name", "age").Find(&users) //SELECT name, age FROM users;
	if result.RowsAffected == 0 {
		fmt.Println("查询失败:暂无记录")
	}
	return users
}
