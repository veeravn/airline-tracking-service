package controllers

import (
	"airline-tracking-service/services"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ✅ Mock Flight Service that implements `FlightServiceInterface`
type MockFlightService struct {
	mock.Mock
}

// ✅ Mock `SearchFlights()` for flight search tests
func (m *MockFlightService) SearchFlights(params map[string]string) ([]services.FlightData, error) {
	args := m.Called(params)
	return args.Get(0).([]services.FlightData), args.Error(1)
}

// ✅ Mock `FetchLiveFlightsWithLocation()` to match `FlightServiceInterface`
func (m *MockFlightService) FetchLiveFlightsWithLocation() ([]services.FlightData, error) {
	args := m.Called()
	return args.Get(0).([]services.FlightData), args.Error(1)
}

func TestSearchFlightsHandler(t *testing.T) {
	mockService := new(MockFlightService)
	flightController := NewFlightController(mockService)

	// ✅ Mock flight data
	mockFlights := []services.FlightData{
		{
			FlightStatus: "active",
			Departure:    services.FlightData{}.Departure,
			Arrival:      services.FlightData{}.Arrival,
			Airline:      services.FlightData{}.Airline,
			Flight:       services.FlightData{}.Flight,
		},
	}

	// ✅ Set test data
	mockFlights[0].Departure.Airport = "Soekarno-Hatta International"
	mockFlights[0].Departure.IATA = "CGK"
	mockFlights[0].Arrival.Airport = "Babullah"
	mockFlights[0].Arrival.IATA = "TTE"
	mockFlights[0].Airline.Name = "Batik Air"
	mockFlights[0].Airline.IATA = "ID"
	mockFlights[0].Flight.Number = "6140"
	mockFlights[0].Flight.IATA = "ID6140"

	mockService.On("SearchFlights", mock.Anything).Return(mockFlights, nil)

	req := httptest.NewRequest("GET", "/api/v1/search-flights?flight_iata=ID6140", nil)
	w := httptest.NewRecorder()
	flightController.SearchFlightsHandler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []services.FlightData
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 1)
	assert.Equal(t, "ID6140", response[0].Flight.IATA)
}
