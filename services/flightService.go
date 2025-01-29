package services

import (
	"airline-tracking-service/config"
	"airline-tracking-service/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Check Redis before making an API request
func GetCachedFlightData(flightNumber string) (*models.FlightData, error) {
	ctx := context.Background()
	val, err := config.RedisClient.Get(ctx, flightNumber).Result()
	if err == nil {
		var flight models.FlightData
		json.Unmarshal([]byte(val), &flight)
		return &flight, nil
	}
	return nil, err
}

// Store flight data in Redis for caching
func CacheFlightData(flightNumber string, flight models.FlightData) {
	ctx := context.Background()
	data, _ := json.Marshal(flight)
	config.RedisClient.Set(ctx, flightNumber, data, 5*time.Minute) // Cache for 5 minutes
}

// Fetch flight data (using Redis caching)
func FetchLiveFlightData(flightNumber string) (*models.FlightData, error) {
	// Check cache first
	if cachedData, err := GetCachedFlightData(flightNumber); err == nil {
		fmt.Println("Returning cached flight data")
		return cachedData, nil
	}

	// Fetch from external API
	apiKey := os.Getenv("AVIATIONSTACK_API_KEY")
	url := fmt.Sprintf("http://api.aviationstack.com/v1/flights?access_key=%s&flight_number=%s", apiKey, flightNumber)

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

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("no flight data found")
	}

	// Cache response
	CacheFlightData(flightNumber, result.Data[0])
	return &result.Data[0], nil
}

// Fetch paginated flight data with filtering
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
