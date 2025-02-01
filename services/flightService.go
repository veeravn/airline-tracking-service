package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// FlightServiceInterface for dependency injection
type FlightServiceInterface interface {
	FetchLiveFlightsWithLocation() ([]FlightData, error)
}

// FlightService struct that implements FlightServiceInterface
type FlightService struct{}

// Position struct for live flight tracking data
type Position struct {
	Latitude        float64 `json:"latitude,omitempty"`
	Longitude       float64 `json:"longitude,omitempty"`
	Altitude        float64 `json:"altitude,omitempty"`
	Direction       float64 `json:"direction,omitempty"`
	SpeedHorizontal float64 `json:"speed_horizontal,omitempty"`
	SpeedVertical   float64 `json:"speed_vertical,omitempty"`
	IsGround        bool    `json:"is_ground,omitempty"`
}

// FlightData struct to match the API response
type FlightData struct {
	FlightDate   string `json:"flight_date,omitempty"`
	FlightStatus string `json:"flight_status,omitempty"`
	Departure    struct {
		Airport   string `json:"airport,omitempty"`
		IATA      string `json:"iata,omitempty"`
		ICAO      string `json:"icao,omitempty"`
		Terminal  string `json:"terminal,omitempty"`
		Gate      string `json:"gate,omitempty"`
		Delay     int    `json:"delay,omitempty"`
		Scheduled string `json:"scheduled,omitempty"`
		Estimated string `json:"estimated,omitempty"`
		Actual    string `json:"actual,omitempty"`
	} `json:"departure,omitempty"`
	Arrival struct {
		Airport   string `json:"airport,omitempty"`
		IATA      string `json:"iata,omitempty"`
		ICAO      string `json:"icao,omitempty"`
		Terminal  string `json:"terminal,omitempty"`
		Gate      string `json:"gate,omitempty"`
		Delay     int    `json:"delay,omitempty"`
		Scheduled string `json:"scheduled,omitempty"`
		Estimated string `json:"estimated,omitempty"`
		Actual    string `json:"actual,omitempty"`
	} `json:"arrival,omitempty"`
	Airline struct {
		Name string `json:"name,omitempty"`
		IATA string `json:"iata,omitempty"`
		ICAO string `json:"icao,omitempty"`
	} `json:"airline,omitempty"`
	Flight struct {
		Number string `json:"number,omitempty"`
		IATA   string `json:"iata,omitempty"`
		ICAO   string `json:"icao,omitempty"`
	} `json:"flight,omitempty"`
	Aircraft struct {
		Registration string `json:"registration,omitempty"`
		IATA         string `json:"iata,omitempty"`
		ICAO         string `json:"icao,omitempty"`
		ICAO24       string `json:"icao24,omitempty"`
	} `json:"aircraft,omitempty"`
	Live Position `json:"live,omitempty"`
}

// APIResponse struct to handle pagination
type APIResponse struct {
	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Count  int `json:"count"`
		Total  int `json:"total"`
	} `json:"pagination"`
	Data []FlightData `json:"data"`
}

// FetchLiveFlightsWithLocation fetches real-time flight data from AviationStack API
func (fs FlightService) FetchLiveFlightsWithLocation() ([]FlightData, error) {
	apiKey := os.Getenv("AVIATIONSTACK_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("API key is missing")
	}

	url := fmt.Sprintf("http://api.aviationstack.com/v1/flights?access_key=%s&limit=100&flight_status=active", apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Print the raw response for debugging
	fmt.Println("API Response:", string(body))

	// Unmarshal into APIResponse struct
	var result APIResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	// Filter out invalid flight data
	var validFlights []FlightData
	for _, flight := range result.Data {
		if flight.Flight.Number != "" && flight.FlightStatus != "" &&
			flight.Live.Latitude != 0 && flight.Live.Longitude != 0 {
			validFlights = append(validFlights, flight)
		}
	}

	// Return only valid flight data
	return validFlights, nil
}
