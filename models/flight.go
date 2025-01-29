package models

type FlightData struct {
	FlightNumber string  `json:"flight_number"`
	Airline      string  `json:"airline"`
	Departure    string  `json:"departure"`
	Arrival      string  `json:"arrival"`
	Status       string  `json:"status"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

type FlightSearchParams struct {
	Airline   string
	Departure string
	Arrival   string
	Date      string
}
