package database

import (
	"30-days-of-robotics-backend/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:2000)/robotics"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Db connection  error: %v", err.Error())
	}
}

func AutoMigrate() {
	err := DB.AutoMigrate(models.User{}, models.Task{}, models.Track{})
	if err != nil {
		log.Printf("Error connecting to Database: %v", err.Error())
	}
}
