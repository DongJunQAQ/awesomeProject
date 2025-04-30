package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model         //嵌入预定义的结构体
	Name       string  `gorm:"not null"` //设置字段非空
	Email      *string // 指针表示该字段可以为空
	Age        uint8   `gorm:"default:0"` //设置字段默认值
	Birthday   *time.Time
}
