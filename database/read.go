package database

import (
	"awesomeProject/model"
	"fmt"
	"gorm.io/gorm"
)

func QuerySingleWithField(db *gorm.DB) model.User { //根据字段条件查询单条记录
	user := model.User{}
	//result := db.First(&user, "name = ?", "dongjun") //根据指定的字段来检索，SELECT * FROM users WHERE name = "dj" ORDER BY id LIMIT 1;
	result := db.Where("name = ?", "dongjun").First(&user)
	if result.Error != nil {
		fmt.Println("查询失败:", result.Error)
	}
	return user
}

func QueryManyWithField(db *gorm.DB) []model.User { //根据字段条件查询多条记录
	var users []model.User
	db.Where("name <> ?", "dj").Find(&users) //SELECT * FROM users WHERE name <> "dongjun";<>为不等于
	if len(users) == 0 {
		fmt.Println("查询失败:暂无记录")
	}
	return users
}

func InQuery(db *gorm.DB) []model.User { //in查询
	var users []model.User
	db.Where("name IN ?", []string{"dongjun", "dj"}).Find(&users) //查询名字为dongjun或dj的数据，SELECT * FROM users WHERE name IN ('dongjun','dj');
	if len(users) == 0 {
		fmt.Println("查询失败:暂无记录")
	}
	return users
}

func LikeQuery(db *gorm.DB) []model.User { //like查询
	var users []model.User
	db.Where("name LIKE ?", "%dj%").Find(&users) //SELECT * FROM users WHERE name LIKE '%dj%';
	if len(users) == 0 {
		fmt.Println("查询失败:暂无记录")
	}
	return users
}

func AndQuery(db *gorm.DB) []model.User { //and查询
	var users []model.User
	db.Where("name = ? AND age >= ?", "dj", "22").Find(&users) //SELECT * FROM users WHERE name = 'dj' AND age >= 22;
	if len(users) == 0 {
		fmt.Println("查询失败:暂无记录")
	}
	return users
}

func BetweenQuery(db *gorm.DB) []model.User { //between...and...查询
	var users []model.User
	db.Where("age BETWEEN ? AND ?", 18, 22).Find(&users) //SELECT * FROM users WHERE age BETWEEN '18' AND '22';
	if len(users) == 0 {
		fmt.Println("查询失败:暂无记录")
	}
	return users
}

func NotQuery(db *gorm.DB) []model.User { //not查询
	var users []model.User
	db.Not([]uint64{29, 30}).Find(&users) //SELECT * FROM users WHERE id NOT IN (29,30);
	if len(users) == 0 {
		fmt.Println("查询失败:暂无记录")
	}
	return users
}

func OrQuery(db *gorm.DB) []model.User { //or查询
	var users []model.User
	db.Where("name = ?", "dj").Or("age = ?", "18").Find(&users) //SELECT * FROM users WHERE name = 'dj' OR age = '18';
	if len(users) == 0 {
		fmt.Println("查询失败:暂无记录")
	}
	return users
}
