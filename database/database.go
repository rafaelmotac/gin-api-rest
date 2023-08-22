package database

import (
	"api-go-gin/models"
	"api-go-gin/properties"
	"api-go-gin/util"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DbConnect() {

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		properties.Properties.Database.Host,
		properties.Properties.Database.User,
		properties.Properties.Database.Password,
		properties.Properties.Database.DBName,
		properties.Properties.Database.Port,
		properties.Properties.Database.SSLMode)

	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	util.LogError(err, "Could not connect to the database")

	err = DB.AutoMigrate(&models.Student{}, &models.User{})

	util.LogError(err, "Could not migrate the database")
}
