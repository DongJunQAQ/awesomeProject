package database

import (
	"awesomeProject/model"
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
	db.Model(&model.User{}).Select("id", "name", "email", "age", "birthday").Find(&users)
	//这里的birthday字段查询结果为nil，即使该字段在数据库中已有记录
	return users
}
