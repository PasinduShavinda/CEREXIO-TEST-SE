package routes

import (
	"github.com/gorilla/mux"
	"github.com/PasinduShavinda/go-project-api/internal/handlers"
	"github.com/PasinduShavinda/go-project-api/internal/middleware"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Authentication
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	// Inventory Master
	router.HandleFunc("/inventory-master", middleware.JwtMiddleware(handlers.GetAllInventoryMaster)).Methods("GET")
	router.HandleFunc("/inventory-master/{id}", middleware.JwtMiddleware(handlers.GetInventoryMasterById)).Methods("GET")
	router.HandleFunc("/inventory-master", middleware.JwtMiddleware(handlers.CreateInventoryMaster)).Methods("POST")
	router.HandleFunc("/inventory-master/{id}", middleware.JwtMiddleware(handlers.UpdateInventoryMaster)).Methods("PUT")
	router.HandleFunc("/inventory-master/{id}", middleware.JwtMiddleware(handlers.DeleteInventoryMaster)).Methods("DELETE")

	// Inventory Details
	router.HandleFunc("/inventory-details", middleware.JwtMiddleware(handlers.GetAllInventoryDetails)).Methods("GET")
	router.HandleFunc("/inventory-details/{id}", middleware.JwtMiddleware(handlers.GetInventoryDetailsById)).Methods("GET")
	router.HandleFunc("/inventory-details", middleware.JwtMiddleware(handlers.CreateInventoryDetails)).Methods("POST")
	router.HandleFunc("/inventory-details/{id}", middleware.JwtMiddleware(handlers.UpdateInventoryDetails)).Methods("PUT")
	router.HandleFunc("/inventory-details/{id}", middleware.JwtMiddleware(handlers.DeleteInventoryDetails)).Methods("DELETE")

	// Define API routes - Inventory
	router.HandleFunc("/inventory-details-overall", middleware.JwtMiddleware(handlers.GetOverallInventoryDetails)).Methods("GET")                
	router.HandleFunc("/inventory-details-overall-filter", middleware.JwtMiddleware(handlers.GetFilteredOverallInventoryDetails)).Methods("GET") 

	return router
}
