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

// WebSocket connection handler for live flight updates
func (wh *WebSocketHandler) LiveFlightUpdates(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade failed:", err)
		http.Error(w, "Failed to upgrade WebSocket", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected to WebSocket")

	// Set WebSocket read deadline to detect inactivity
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))

	// Keep connection alive by resetting deadline on each pong message
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		return nil
	})

	for {
		// Fetch live flight data
		flights, err := wh.FlightService.FetchLiveFlightsWithLocation()
		if err != nil {
			fmt.Println("Error fetching live flights:", err)
			break
		}

		// Set write deadline to avoid infinite blocking
		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

		// Send flight data to the WebSocket client
		err = conn.WriteJSON(flights)
		if err != nil {
			fmt.Println("WebSocket Error (sending data):", err)
			break
		}

		// Send a ping to check if client is still connected
		err = conn.WriteMessage(websocket.PingMessage, nil)
		if err != nil {
			fmt.Println("Client disconnected:", err)
			break
		}

		// Wait before sending the next update
		time.Sleep(10 * time.Second)
	}

	fmt.Println("WebSocket connection closed")
}
