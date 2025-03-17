package database

import (
	"awesomeProject/model"
	"fmt"
	"gorm.io/gorm"
)

func QuerySingleWithField(db *gorm.DB) model.User {
	user := model.User{}
	//result := db.First(&user, "name = ?", "dongjun") //根据指定的字段来检索，SELECT * FROM users WHERE name = "dj" ORDER BY id LIMIT 1;
	result := db.Where("name = ?", "dongjun").First(&user)
	if result.Error != nil {
		fmt.Println("查询失败:", result.Error)
	}
	return user
}

func QueryManyWithField(db *gorm.DB) []model.User {
	var users []model.User
	db.Where("name <> ?", "dongjun").Find(&users) //SELECT * FROM users WHERE name <> "dongjun";<>为不等于
	if len(users) == 0 {
		fmt.Println("查询失败:暂无记录")
	}
	return users
}
func InQuery() {}
