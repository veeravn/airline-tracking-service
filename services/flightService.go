package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// FlightData struct
type FlightData struct {
	FlightNumber string  `json:"flight_number"`
	Airline      string  `json:"airline"`
	Departure    string  `json:"departure"`
	Arrival      string  `json:"arrival"`
	Status       string  `json:"status"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

// Define an interface to allow mocking
type FlightServiceInterface interface {
	FetchLiveFlightsWithLocation() ([]FlightData, error)
}

// Real implementation of FlightService
type FlightService struct{}

// FetchLiveFlightsWithLocation fetches real-time flight data
func (fs FlightService) FetchLiveFlightsWithLocation() ([]FlightData, error) {
	apiKey := os.Getenv("AVIATIONSTACK_API_KEY")
	if apiKey == "" {
		log.Println("API key is missing.")
		return nil, fmt.Errorf("API key is missing")
	}

	url := fmt.Sprintf("http://api.aviationstack.com/v1/flights?access_key=%s", apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error making request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Data []FlightData `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Data, nil
}
