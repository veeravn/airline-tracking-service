package controllers

import (
	"airline-tracking-service/services"
	"encoding/json"
	"fmt"
	"net/http"
)

// FlightController struct to inject FlightService
type FlightController struct {
	FlightService services.FlightServiceInterface
}

// Constructor for FlightController
func NewFlightController(flightService services.FlightServiceInterface) *FlightController {
	return &FlightController{FlightService: flightService}
}

// LiveFlightsHandler handles requests for real-time flights
func (fc *FlightController) LiveFlightsHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch live flight data
	flights, err := fc.FlightService.FetchLiveFlightsWithLocation()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching live flights: %s", err), http.StatusInternalServerError)
		fmt.Println("Error fetching live flights:", err)
		return
	}

	// If no flights found, return a meaningful response
	if len(flights) == 0 {
		http.Error(w, "No live flights available", http.StatusNotFound)
		fmt.Println("No live flights found")
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Send the flight data as JSON response
	if err := json.NewEncoder(w).Encode(flights); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %s", err), http.StatusInternalServerError)
		fmt.Println("Error encoding JSON response:", err)
	}
}
