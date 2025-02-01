package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockFlightService implements FlightServiceInterface for testing
type MockFlightService struct{}

// Mock response data
func (m MockFlightService) FetchLiveFlightsWithLocation() ([]FlightData, error) {
	return []FlightData{
		{
			FlightNumber: "AA123",
			Airline:      "American Airlines",
			Departure:    "JFK",
			Arrival:      "LAX",
			Status:       "On Time",
			Latitude:     40.6413,
			Longitude:    -73.7781,
		},
		{
			FlightNumber: "BA456",
			Airline:      "British Airways",
			Departure:    "LHR",
			Arrival:      "JFK",
			Status:       "Delayed",
			Latitude:     51.4700,
			Longitude:    -0.4543,
		},
	}, nil
}

// Test fetching live flight data from the mock service
func TestFetchLiveFlightsWithLocation(t *testing.T) {
	// Use the mock service instead of the real one
	mockService := MockFlightService{}

	// Fetch flight data
	flights, err := mockService.FetchLiveFlightsWithLocation()
	assert.NoError(t, err)
	assert.NotEmpty(t, flights)

	// Validate data
	assert.Equal(t, "AA123", flights[0].FlightNumber)
	assert.Equal(t, "JFK", flights[0].Departure)
	assert.Equal(t, "On Time", flights[0].Status)

	assert.Equal(t, "BA456", flights[1].FlightNumber)
	assert.Equal(t, "LHR", flights[1].Departure)
	assert.Equal(t, "Delayed", flights[1].Status)
}
