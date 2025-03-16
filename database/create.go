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

func CreateManyRecord(db *gorm.DB) int64 { //添加多条记录
	users := []*model.User{ //通过切片创建多条记录
		{Name: "Jinzhu", Age: 18},
		{Name: "Jackson", Age: 19},
	}
	result := db.Create(&users)
	return result.RowsAffected
}

func SelectedFields(db *gorm.DB) int64 { //添加带有指定字段的记录
	newUser := model.User{Name: "dj", Age: 20}
	result := db.Select("Name", "Age").Create(&newUser) //INSERT INTO `users` (`name`,`age`) VALUES ("dj", 20);
	return result.RowsAffected
}

func IgnoreFields(db *gorm.DB) int64 { //创建记录时忽略指定的字段
	email := "2360052818@qq.com"
	newUser := model.User{Name: "dongjun", Age: 18, Email: &email}
	result := db.Omit("Name", "Age").Create(&newUser) //只会插入Email字段
	return result.RowsAffected
}
