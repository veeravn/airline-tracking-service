package main

import (
	"airline-tracking-service/config"
	"airline-tracking-service/controllers"
	"airline-tracking-service/services"
	"fmt"
	"net/http"
)

// Enable CORS Middleware
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow frontend to access API
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	config.ConnectRedis()

	flightService := services.FlightService{}
	flightController := controllers.NewFlightController(flightService)
	webSocketHandler := controllers.NewWebSocketHandler(flightService)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/live-flights", flightController.LiveFlightsHandler)
	mux.HandleFunc("/api/v1/search-flights", flightController.SearchFlightsHandler)
	mux.HandleFunc("/ws/live-updates", webSocketHandler.LiveFlightUpdates)

	// Apply CORS middleware
	handler := enableCORS(mux)

	port := "8080"
	fmt.Println("Server running on http://localhost:" + port)
	http.ListenAndServe(":"+port, handler)
}
