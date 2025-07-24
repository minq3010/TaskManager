package database

import (
	"gorm.io/gorm"
	"TaskManager/models"
	"log"
	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Can not connect to database", err)
	}

	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("Error when migrate Database", err)
	}
}
