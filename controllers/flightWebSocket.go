package controllers

import (
	"airline-tracking-service/services"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// WebSocket handler for live flight updates
func LiveFlightUpdates(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	for {
		flightData, err := services.FetchLiveFlights("", "", "", "", 1, 5)
		if err != nil {
			fmt.Println("Error fetching live flights:", err)
			break
		}

		// Send updates to frontend
		conn.WriteJSON(flightData)
		time.Sleep(10 * time.Second) // Update every 10 seconds
	}
}
