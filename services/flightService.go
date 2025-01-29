package services

import (
	"airline-tracking-service/config"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// FlightData includes flight details along with location for mapping
type FlightData struct {
	FlightNumber string  `json:"flight_number"`
	Airline      string  `json:"airline"`
	Departure    string  `json:"departure"`
	Arrival      string  `json:"arrival"`
	Status       string  `json:"status"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

// Check Redis before making an API request
func GetCachedFlightData(flightNumber string) (*FlightData, error) {
	ctx := context.Background()
	val, err := config.RedisClient.Get(ctx, flightNumber).Result()
	if err == nil {
		var flight FlightData
		json.Unmarshal([]byte(val), &flight)
		return &flight, nil
	}
	return nil, err
}

// Store flight data in Redis for caching
func CacheFlightData(flightNumber string, flight FlightData) {
	ctx := context.Background()
	data, _ := json.Marshal(flight)
	config.RedisClient.Set(ctx, flightNumber, data, 5*time.Minute) // Cache for 5 minutes
}

// Fetch live flight data from AviationStack API
func FetchLiveFlightsWithLocation() ([]FlightData, error) {
	apiKey := os.Getenv("AVIATIONSTACK_API_KEY")
	if apiKey == "" {
		log.Println("API key is missing. Check .env file.")
		return nil, fmt.Errorf("API key is missing")
	}

	url := fmt.Sprintf("http://api.aviationstack.com/v1/flights?access_key=%s", apiKey)

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
		Data []FlightData `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	// Filter out flights with missing location data
	var flightsWithLocation []FlightData
	for _, flight := range result.Data {
		if flight.Latitude != 0 && flight.Longitude != 0 {
			flightsWithLocation = append(flightsWithLocation, flight)
		}
	}

	return flightsWithLocation, nil
}

// Fetch flights with pagination and filtering
func FetchFilteredFlights(flightNumber, airline, departure, arrival string, page, pageSize int) ([]FlightData, error) {
	allFlights, err := FetchLiveFlightsWithLocation()
	if err != nil {
		return nil, err
	}

	// Apply filters
	var filteredFlights []FlightData
	for _, flight := range allFlights {
		if flightNumber != "" && flight.FlightNumber != flightNumber {
			continue
		}
		if airline != "" && flight.Airline != airline {
			continue
		}
		if departure != "" && flight.Departure != departure {
			continue
		}
		if arrival != "" && flight.Arrival != arrival {
			continue
		}
		filteredFlights = append(filteredFlights, flight)
	}

	// Implement pagination
	start := (page - 1) * pageSize
	end := start + pageSize
	if start > len(filteredFlights) {
		return []FlightData{}, nil
	}
	if end > len(filteredFlights) {
		end = len(filteredFlights)
	}

	return filteredFlights[start:end], nil
}
