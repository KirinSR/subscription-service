package main

import (
	"log"
	"net/http"

	"kirin.com/subscription-service/config"
	"kirin.com/subscription-service/routes"
	"kirin.com/subscription-service/utils"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize database
	db := utils.InitDB(config.GetDBConfig())

	// Initialize router
	router := utils.InitRouter()

	// Set up routes
	routes.RegisterRoutes(router, db)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))
}
