package model

// migration 自动迁移模式
func migration() {
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}).
		AutoMigrate(&Task{})
	// 外键关联
	DB.Model(&Task{}).AddForeignKey("Uid", "User(id)", "CASCADE", "CASCADE")
}
