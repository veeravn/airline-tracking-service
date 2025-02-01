package controllers

import (
	"airline-tracking-service/services"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocket upgrader settings
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // Allow all origins (modify for security)
}

// WebSocket handler for live flight updates
type WebSocketHandler struct {
	FlightService services.FlightServiceInterface
}

// Constructor function for WebSocketHandler
func NewWebSocketHandler(flightService services.FlightServiceInterface) *WebSocketHandler {
	return &WebSocketHandler{FlightService: flightService}
}

// Handle WebSocket connections for real-time updates
func (wh *WebSocketHandler) LiveFlightUpdates(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	// Continuously send flight updates to the client
	for {
		flights, err := wh.FlightService.FetchLiveFlightsWithLocation()
		if err != nil {
			fmt.Println("Error fetching live flights:", err)
			break
		}

		// Send flight data to the WebSocket client
		if err := conn.WriteJSON(flights); err != nil {
			fmt.Println("Error sending data over WebSocket:", err)
			break
		}

		time.Sleep(10 * time.Second) // Update clients every 10 seconds
	}
}
