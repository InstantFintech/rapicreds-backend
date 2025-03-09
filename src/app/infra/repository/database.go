package repository

import (
	"rapicreds-backend/src/app/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:password123@tcp(localhost:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.UserLoan{})
	return db
}
