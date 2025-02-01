package controllers

import (
	"airline-tracking-service/services"
	"encoding/json"
	"fmt"
	"net/http"
)

// Create a variable to hold the flight service instance
var flightService services.FlightServiceInterface = services.FlightService{}

// Function to set a custom flight service (Used for testing)
func SetFlightService(service services.FlightServiceInterface) {
	flightService = service
}

// LiveFlightsHandler uses the injected service
func LiveFlightsHandler(w http.ResponseWriter, r *http.Request) {
	flights, err := flightService.FetchLiveFlightsWithLocation()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching live flights: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flights)
}
