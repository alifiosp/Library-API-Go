package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModelsbooks() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(localhost)/booksapigo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("connection failed")
	}

	db.AutoMigrate(&Books{})

	return db
}
