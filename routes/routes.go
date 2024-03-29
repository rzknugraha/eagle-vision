// Package routes consist of router path used for handling incoming request //
package routes

import (
	"github.com/gorilla/mux"
	"github.com/rzknugraha/eagle-vision/controllers"
)

// Route is
type Route struct{}

// Init is
func (r *Route) Init() *mux.Router {
	// Initialize controller //
	healthCheckController := controllers.InitHealthCheckController()
	playerController := controllers.InitPlayerController()

	// Initialize router //
	router := mux.NewRouter().StrictSlash(false)
	v1 := router.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/healthcheck", healthCheckController.HealthCheck).Methods("GET")
	v1.HandleFunc("/player", playerController.StorePlayer).Methods("POST")

	return v1
}
