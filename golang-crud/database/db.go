package database

import (
	"golang-crud/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// mysql
	dsn := "root:@tcp(localhost:3306)/db_simple_crud?parseTime=true"
	var err error
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.Post{})

	DB = database
}
