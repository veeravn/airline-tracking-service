package controllers

import (
	"airline-tracking-service/services"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock WebSocket Service
type MockWebSocketService struct {
	mock.Mock
}

func (m *MockWebSocketService) FetchLiveFlightsWithLocation() ([]services.FlightData, error) {
	args := m.Called()
	return args.Get(0).([]services.FlightData), args.Error(1)
}

// Mocking SearchFlights (needed for FlightServiceInterface compatibility)
func (m *MockWebSocketService) SearchFlights(params map[string]string) ([]services.FlightData, error) {
	args := m.Called(params)
	return args.Get(0).([]services.FlightData), args.Error(1)
}

func TestLiveFlightUpdates(t *testing.T) {
	mockService := new(MockWebSocketService)
	webSocketHandler := NewWebSocketHandler(mockService)

	// ✅ Ensure `Departure` and `Arrival` match `FlightData`'s internal struct definitions
	mockFlights := []services.FlightData{
		{
			FlightStatus: "active",
			Departure:    services.FlightData{}.Departure, // ✅ Use FlightData's Departure struct
			Arrival:      services.FlightData{}.Arrival,   // ✅ Use FlightData's Arrival struct
			Airline:      services.FlightData{}.Airline,   // ✅ Use FlightData's Airline struct
			Flight:       services.FlightData{}.Flight,    // ✅ Use FlightData's Flight struct
			Live: services.Position{
				Latitude: -4.9122, Longitude: 108.64, Altitude: 10210.8,
			},
		},
	}

	// ✅ Assign data to nested structs
	mockFlights[0].Departure.Airport = "Soekarno-Hatta International"
	mockFlights[0].Departure.IATA = "CGK"
	mockFlights[0].Arrival.Airport = "Babullah"
	mockFlights[0].Arrival.IATA = "TTE"
	mockFlights[0].Airline.Name = "Batik Air"
	mockFlights[0].Airline.IATA = "ID"
	mockFlights[0].Flight.Number = "6140"
	mockFlights[0].Flight.IATA = "ID6140"

	mockService.On("FetchLiveFlightsWithLocation").Return(mockFlights, nil)

	server := httptest.NewServer(http.HandlerFunc(webSocketHandler.LiveFlightUpdates))
	defer server.Close()

	url := "ws" + server.URL[4:]
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	assert.NoError(t, err)
	defer conn.Close()

	time.Sleep(time.Second * 2) // Allow data to be sent

	var receivedFlights []services.FlightData
	err = conn.ReadJSON(&receivedFlights)
	assert.NoError(t, err)
	assert.Len(t, receivedFlights, 1)
	assert.Equal(t, "ID6140", receivedFlights[0].Flight.IATA)
}
