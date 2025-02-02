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

// SearchFlightsHandler processes flight search requests
func (fc *FlightController) SearchFlightsHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	params := map[string]string{
		"flight_iata":   r.URL.Query().Get("flight_iata"),
		"airline_iata":  r.URL.Query().Get("airline_iata"),
		"dep_iata":      r.URL.Query().Get("dep_iata"),
		"arr_iata":      r.URL.Query().Get("arr_iata"),
		"flight_status": r.URL.Query().Get("flight_status"),
	}

	// Fetch filtered flights
	flights, err := fc.FlightService.SearchFlights(params)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching flights: %s", err), http.StatusInternalServerError)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flights)
}
