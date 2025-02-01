package controllers

import (
	"airline-tracking-service/services"
	"encoding/json"
	"fmt"
	"net/http"
)

// LiveFlightsHandler handles requests for real-time flights with location
func LiveFlightsHandler(w http.ResponseWriter, r *http.Request) {
	flights, err := services.FetchLiveFlightsWithLocation()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching live flights: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flights)
}
