package main

import (
	"api-go-gin/config"
	"api-go-gin/database"
	"api-go-gin/properties"
	"api-go-gin/routes"
)

// @title Students API
// @version 1.0
// @description This is a sample server for managing students.
// @host localhost:9000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	properties.InitProperties()
	config.InitValidator()
	database.DbConnect()
	routes.HandleRequests()
}
