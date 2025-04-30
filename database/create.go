package database

import (
	"awesomeProject/model"
	"gorm.io/gorm"
	"time"
)

func CreateOneRecord(db *gorm.DB) int64 { //添加一条记录
	birthday, _ := time.Parse("2006-01-02", "2002-10-31")
	newUser := model.User{Name: "DongJun", Age: 18, Birthday: &birthday} //因为在模型中定义的Birthday字段值为指针类型，所以这里传参时需要取地址
	result := db.Create(&newUser)
	return result.RowsAffected //返回已插入记录的条数
}
