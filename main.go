package main

import (
	"airline-tracking-service/config"
	"airline-tracking-service/controllers"
	"airline-tracking-service/utils"
	"fmt"
	"net/http"
)

func main() {
	// Connect to Redis
	config.ConnectRedis()
	config.LoadEnv()
	utils.SetupLogger()

	// API Endpoints
	http.HandleFunc("/api/v1/flights", controllers.SearchFlightsHandler)
	http.HandleFunc("/ws/live-updates", controllers.LiveFlightUpdates)

	port := config.GetConfig("PORT", "8080")
	utils.Logger.Info(fmt.Sprintf("Server running on http://localhost:%s", port))
	http.ListenAndServe(":"+port, nil)
}
