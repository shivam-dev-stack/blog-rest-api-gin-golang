package config

import (
	"gin-blog-api/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	err = database.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("Migration failed!", err)
	}

	database.AutoMigrate(&models.Post{}, &models.User{})

	DB = database
}
