package main

import (
	"api-go-gin/database"
	"api-go-gin/routes"
	"api-go-gin/services"
)

func main() {
	services.InitValidator()
	database.DbConnect()
	routes.HandleRequests()
}
