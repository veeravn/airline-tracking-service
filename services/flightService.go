package services

import (
	"airline-tracking-service/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// FetchLiveFlightData fetches real-time flight data from AviationStack API
func FetchLiveFlights(flightNumber, airline, departure, arrival string, page, pageSize int) ([]models.FlightData, error) {
	apiKey := os.Getenv("AVIATIONSTACK_API_KEY")
	if apiKey == "" {
		log.Println("API key is missing. Check .env file.")
		return nil, fmt.Errorf("API key is missing")
	}

	url := fmt.Sprintf("http://api.aviationstack.com/v1/flights?access_key=%s&page=%d&limit=%d", apiKey, page, pageSize)

	// Append filters if provided
	if flightNumber != "" {
		url += fmt.Sprintf("&flight_number=%s", flightNumber)
	}
	if airline != "" {
		url += fmt.Sprintf("&airline_name=%s", airline)
	}
	if departure != "" {
		url += fmt.Sprintf("&departure_iata=%s", departure)
	}
	if arrival != "" {
		url += fmt.Sprintf("&arrival_iata=%s", arrival)
	}

	log.Println("Fetching URL:", url)

	// Make HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error making request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	var result struct {
		Data []models.FlightData `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Data, nil
}
