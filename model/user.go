package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model //嵌入预定义的结构体
	//ID           uint // 主键，gorm.Model结构体中自带ID主键字段
	Name         string  `gorm:"not null"` //设置字段非空
	Email        *string // 指针表示该字段可以为空
	Age          uint8   `gorm:"default:0"` //设置字段默认值
	Birthday     *time.Time
	MemberNumber sql.NullString // 使用sql.NullString处理可空字符串
	ActivatedAt  sql.NullTime   // 使用sql.NullTime作为可空时间字段
}
