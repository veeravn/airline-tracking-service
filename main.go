package main

import (
	"airline-tracking-service/config"
	"airline-tracking-service/controllers"
	"airline-tracking-service/services"
	"fmt"
	"net/http"
)

func main() {
	// Connect to Redis
	config.ConnectRedis()

	// Create the flight service and controllers
	flightService := services.FlightService{}
	flightController := controllers.NewFlightController(flightService)
	webSocketHandler := controllers.NewWebSocketHandler(flightService)

	// API Endpoints
	http.HandleFunc("/api/v1/live-flights", flightController.LiveFlightsHandler)
	http.HandleFunc("/ws/live-updates", webSocketHandler.LiveFlightUpdates)

	port := "8080"
	fmt.Println("Server running on http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
