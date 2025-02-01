package controllers

import (
	"airline-tracking-service/services"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockFlightService implements FlightServiceInterface for testing
type MockFlightService struct{}

// Mock response data
func (m MockFlightService) FetchLiveFlightsWithLocation() ([]services.FlightData, error) {
	return []services.FlightData{
		{
			FlightNumber: "AA123",
			Airline:      "American Airlines",
			Departure:    "JFK",
			Arrival:      "LAX",
			Status:       "On Time",
			Latitude:     40.6413,
			Longitude:    -73.7781,
		},
	}, nil
}

// Test the LiveFlightsHandler API response
func TestLiveFlightsHandler(t *testing.T) {
	// Inject the mock service into the controller
	SetFlightService(MockFlightService{})

	// Create a new HTTP request
	req, _ := http.NewRequest("GET", "/api/v1/live-flights", nil)
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(LiveFlightsHandler)
	handler.ServeHTTP(rr, req)

	// Check HTTP status
	assert.Equal(t, http.StatusOK, rr.Code)

	// Parse response body
	var flights []services.FlightData
	json.Unmarshal(rr.Body.Bytes(), &flights)

	// Validate response data
	assert.Equal(t, "AA123", flights[0].FlightNumber)
	assert.Equal(t, "JFK", flights[0].Departure)
}
