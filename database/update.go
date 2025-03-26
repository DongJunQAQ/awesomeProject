package database

import (
	"awesomeProject/model"
	"fmt"
	"gorm.io/gorm"
)

func UpdateSingleColumn(db *gorm.DB) int64 { //更新单列数据，UPDATE users SET name='hello' WHERE id=29;
	result := db.Model(&model.User{}).Where("id = ?", 29).Update("name", "hello12")
	if result.RowsAffected == 0 {
		fmt.Println("无数据更新")
	}
	return result.RowsAffected
}

func UpdatesMultipleColumn(db *gorm.DB) int64 { //更新多列数据，UPDATE users SET name='hello444',age=100 WHERE id=29;
	result := db.Model(&model.User{}).Where("id = ?", 29).Updates(map[string]interface{}{"name": "hello444", "age": 100}) //使用结构体更新时GORM将仅更新非零字段,因此需要使用map来更新属性
	if result.RowsAffected == 0 {
		fmt.Println("无数据更新")
	}
	return result.RowsAffected
}

func BatchUpdates(db *gorm.DB) int64 { //批量更新多条数据，UPDATE users SET age=21;
	result := db.Exec("UPDATE users SET age=21;") //GORM默认会阻止全局更新并且会返回WHERE conditions required错误，因此需要使用原始SQL来更新数据
	if result.RowsAffected == 0 {
		fmt.Println("无数据更新")
	}
	return result.RowsAffected
}
