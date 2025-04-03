package main

import (
	"log"
	"net/http"
	"github.com/PasinduShavinda/go-project-api/internal/routes"
	"github.com/PasinduShavinda/go-project-api/internal/db"
	_ "github.com/PasinduShavinda/go-project-api/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Inventory Management API
// @version 1.0
// @description This is an API for managing inventory
// @host localhost:8000
// @BasePath /
func main() {
	// Initialize database connection
	db.ConnectDB()

	// Initialize router
	router := routes.SetupRouter()

	// Serve Swagger documentation
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start server
	log.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
