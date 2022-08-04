package model

import "github.com/jinzhu/gorm"

// User 用户表
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PasswordDigest string // 存储的是密文
}
