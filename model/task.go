package model

import "github.com/jinzhu/gorm"

// Task 备忘录表
type Task struct {
	gorm.Model
	User      User   `gorm:"ForeignKey:Uid"` // 该条Task属于哪个用户
	Uid       uint   `gorm:"Not null"`
	Title     string `gorm:"Not null"`
	Status    int    `gorm:"default:0"` // 0: Task未完成，1：Task已完成
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64
}
