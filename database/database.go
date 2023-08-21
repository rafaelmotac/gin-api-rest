package database

import (
	"api-go-gin/models"
	"api-go-gin/util"
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

	util.LogError(err, "Could not connect to the database")

	err = DB.AutoMigrate(&models.Student{}, &models.User{})

	util.LogError(err, "Could not migrate the database")
}
