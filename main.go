package main

import (
	"airline-tracking-service/config"
	"airline-tracking-service/controllers"
	"fmt"
	"net/http"
)

func main() {
	// Connect to Redis
	config.ConnectRedis()

	// API Endpoints
	http.HandleFunc("/api/v1/live-flights", controllers.LiveFlightsHandler)
	http.HandleFunc("/ws/live-updates", controllers.LiveFlightUpdates) // WebSocket route

	port := "8080"
	fmt.Println("Server running on http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
