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

// Fetch live flight data from API
func FetchLiveFlightsWithLocation() ([]FlightData, error) {
	apiKey := os.Getenv("AVIATIONSTACK_API_KEY")
	url := fmt.Sprintf("http://api.aviationstack.com/v1/flights?access_key=%s", apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching flight data:", err)
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
