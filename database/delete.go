package database

import (
	"awesomeProject/model"
	"fmt"
	"gorm.io/gorm"
)

func DeleteSingleWithPrimaryKey(db *gorm.DB) int64 { //按主键删除单条记录，DELETE from users where id = 31;
	result := db.Delete(&model.User{}, 31)
	if result.RowsAffected == 0 {
		fmt.Println("无数据删除")
	}
	return result.RowsAffected
}

func DeleteMultipleWithPrimaryKey(db *gorm.DB) int64 { //按主键删除多条记录，DELETE FROM users WHERE id IN (34,35);
	result := db.Delete(&model.User{}, []int{34, 35})
	if result.RowsAffected == 0 {
		fmt.Println("无数据删除")
	}
	return result.RowsAffected
}

func DeleteWithField(db *gorm.DB) int64 { //按条件删除记录，DELETE from users where name = "jinzhu";实际上是进行软删除即以修改deleted_at的时间代表删除(需使用预定义的gorm.Model结构体)
	result := db.Where("name = ?", "jinzhu").Delete(&model.User{})
	if result.RowsAffected == 0 {
		fmt.Println("无数据删除")
	}
	return result.RowsAffected
}

func DeleteWithLike(db *gorm.DB) int64 { //按LIKE条件删除记录，DELETE from users where name LIKE "%dj%";
	result := db.Where("name LIKE ?", "%dj%").Delete(&model.User{})
	if result.RowsAffected == 0 {
		fmt.Println("无数据删除")
	}
	return result.RowsAffected
}

func RealDelete(db *gorm.DB) int64 { //真实删除删除记录（非软删除）
	result := db.Unscoped().Where("name = ?", "jinzhu").Delete(&model.User{})
	if result.RowsAffected == 0 {
		fmt.Println("无数据删除")
	}
	return result.RowsAffected
}
