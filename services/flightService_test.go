package services

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock API response for testing
const mockAPIResponse = `{
  "data": [
    {
      "flight_date": "2025-02-02",
      "flight_status": "active",
      "departure": {
        "airport": "Soekarno-Hatta International",
        "iata": "CGK"
      },
      "arrival": {
        "airport": "Babullah",
        "iata": "TTE"
      },
      "airline": {
        "name": "Batik Air",
        "iata": "ID"
      },
      "flight": {
        "number": "6140",
        "iata": "ID6140"
      },
      "live": {
        "latitude": -4.9122,
        "longitude": 108.64,
        "altitude": 10210.8
      }
    }
  ]
}`

// Mock server for testing
func mockServer(response string, statusCode int) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		fmt.Fprintln(w, response)
	})
	return httptest.NewServer(handler)
}

// Test fetching flights with search parameters
func TestSearchFlights(t *testing.T) {
	mockSrv := mockServer(mockAPIResponse, http.StatusOK)
	defer mockSrv.Close()

	// Set environment variable to use mock server URL
	os.Setenv("AVIATIONSTACK_API_KEY", "test_api_key")

	// Override the base API URL
	baseURL := mockSrv.URL

	flightService := FlightService{}

	// Test search by flight number
	t.Run("Search by flight number", func(t *testing.T) {
		params := map[string]string{"flight_iata": "ID6140"}
		flights, err := flightService.SearchFlights(params)
		assert.NoError(t, err)
		assert.Len(t, flights, 1)
		assert.Equal(t, "ID6140", flights[0].Flight.IATA)
	})

	// Test search by airline
	t.Run("Search by airline", func(t *testing.T) {
		params := map[string]string{"airline_iata": "ID"}
		flights, err := flightService.SearchFlights(params)
		assert.NoError(t, err)
		assert.Len(t, flights, 1)
		assert.Equal(t, "Batik Air", flights[0].Airline.Name)
	})

	// Test search by departure airport
	t.Run("Search by departure airport", func(t *testing.T) {
		params := map[string]string{"dep_iata": "CGK"}
		flights, err := flightService.SearchFlights(params)
		assert.NoError(t, err)
		assert.Len(t, flights, 1)
		assert.Equal(t, "CGK", flights[0].Departure.IATA)
	})

	// Test search by arrival airport
	t.Run("Search by arrival airport", func(t *testing.T) {
		params := map[string]string{"arr_iata": "TTE"}
		flights, err := flightService.SearchFlights(params)
		assert.NoError(t, err)
		assert.Len(t, flights, 1)
		assert.Equal(t, "TTE", flights[0].Arrival.IATA)
	})

	// Test search by flight status
	t.Run("Search by flight status", func(t *testing.T) {
		params := map[string]string{"flight_status": "active"}
		flights, err := flightService.SearchFlights(params)
		assert.NoError(t, err)
		assert.Len(t, flights, 1)
		assert.Equal(t, "active", flights[0].FlightStatus)
	})

	// Test API returning an empty response
	t.Run("No flights found", func(t *testing.T) {
		mockSrvEmpty := mockServer(`{"data": []}`, http.StatusOK)
		defer mockSrvEmpty.Close()

		baseURL = mockSrvEmpty.URL
		params := map[string]string{"flight_iata": "UNKNOWN"}
		flights, err := flightService.SearchFlights(params)
		assert.NoError(t, err)
		assert.Len(t, flights, 0)
	})

	// Test API returning an error
	t.Run("API error response", func(t *testing.T) {
		mockSrvError := mockServer(`{"error": "Invalid API Key"}`, http.StatusUnauthorized)
		defer mockSrvError.Close()

		baseURL = mockSrvError.URL
		params := map[string]string{"flight_iata": "ID6140"}
		flights, err := flightService.SearchFlights(params)
		assert.Error(t, err)
		assert.Len(t, flights, 0)
	})
}
