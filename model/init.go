package model

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

// DB 全局数据库操作变量
var DB *gorm.DB

func Database(connstring string) {
	
	db, err := gorm.Open("mysql", connstring)
	if err != nil {
		panic("MySQL数据库连接错误")
	}
	
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	db.SingularTable(true)       // 表名不加s，user表，非users表
	db.DB().SetMaxIdleConns(20)  //设置连接池
	db.DB().SetMaxOpenConns(100) //最大连接数
	db.DB().SetConnMaxIdleTime(time.Second * 30)
	DB = db
}
