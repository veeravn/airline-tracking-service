package main

import (
	"airline-tracking-service/config"
	"airline-tracking-service/controllers"
	"airline-tracking-service/utils"
	"fmt"
	"net/http"
)

func main() {
	// Initialize configuration and logger
	config.LoadEnv()
	utils.SetupLogger()

	// Set up HTTP server and routes
	http.HandleFunc("/api/v1/flights", controllers.SearchFlightsHandler)

	port := config.GetConfig("PORT", "8080")
	utils.Logger.Info(fmt.Sprintf("Server running on http://localhost:%s", port))
	http.ListenAndServe(":"+port, nil)
}
