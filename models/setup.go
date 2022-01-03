package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModelsBooks() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(localhost)/booksapigo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("connection failed")
	}

	db.AutoMigrate(&Books{})

	return db
}

func SetupModelsUsers() *gorm.DB {
	db1, err := gorm.Open("mysql", "root:@(localhost)/usersapigo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("connection failed")
	}

	db1.AutoMigrate(&Users{})

	return db1
}
