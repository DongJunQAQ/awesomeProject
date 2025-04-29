package database

import (
	"awesomeProject/model"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type apiUser struct {
	Id       uint
	Name     string
	Email    string
	Age      uint8
	BirthDay time.Time
}

func SelectFieldsQueryAllUser(db *gorm.DB) []apiUser { //选择指定字段检索返回
	var users []apiUser
	result := db.Model(&model.User{}).Select("id", "name", "email", "age", "birthday").Find(&users)
	if result.RowsAffected == 0 {
		fmt.Println("查询失败:暂无记录")
	}
	return users
}
