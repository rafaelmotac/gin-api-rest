package main

import (
	"api-go-gin/config"
	"api-go-gin/database"
	"api-go-gin/routes"
)

func main() {
	config.InitValidator()
	database.DbConnect()
	routes.HandleRequests()
}
