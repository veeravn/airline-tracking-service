package controllers

import (
	"airline-tracking-service/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func SearchFlightsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	flightNumber := r.URL.Query().Get("flightNumber")
	airline := r.URL.Query().Get("airline")
	departure := r.URL.Query().Get("departure")
	arrival := r.URL.Query().Get("arrival")

	// Pagination parameters
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	// Fetch data from the service
	flights, err := services.FetchFilteredFlights(flightNumber, airline, departure, arrival, page, pageSize)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching live flights: %s", err), http.StatusInternalServerError)
		return
	}

	// Respond with JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flights)
}
