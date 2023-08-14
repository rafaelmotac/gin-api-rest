package database

import (
	"api-go-gin/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DbConnect() {
	connectionString := "host=localhost user=root password=root dbname=root port=30010 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Panic("Could not connect to the database")
	}

	DB.AutoMigrate(&models.Student{})

}
