package database

import (
	"awesomeProject/model"
	"fmt"
	"gorm.io/gorm"
)

func DeleteRecord(db *gorm.DB) int64 { //删除单条记录，会进行软删除即以修改deleted_at的时间代表删除(需使用预定义的gorm.Model结构体)，DELETE from users where name = "jinzhu";
	result := db.Where("name = ?", "jinzhu").Delete(&model.User{})
	if result.RowsAffected == 0 {
		fmt.Println("无数据删除")
	}
	return result.RowsAffected
}
