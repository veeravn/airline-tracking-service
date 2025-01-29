# Airline Tracking Service

## Overview
This project consists of a **real-time flight tracking service** that fetches live flight data using the **AviationStack API**. The system is built with a **Golang backend** and a **React frontend**, supporting **search, pagination, and dynamic filtering**.

---

# Backend Service (Golang)

## Features
- Fetches **real-time flight data** from the **AviationStack API**.
- Supports **search** by flight number, airline, departure, and arrival airports.
- Implements **pagination** for large datasets.
- Exposes a **REST API** for frontend integration.

## Technologies Used
- **Go (Golang)**
- **AviationStack API**
- **net/http** for handling API requests
- **encoding/json** for JSON parsing

## Installation
### Prerequisites
- Go 1.21+
- AviationStack API Key

### Setup
1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/airline-tracking-service.git
   cd airline-tracking-service
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Set up the `.env` file:
   ```plaintext
   AVIATIONSTACK_API_KEY=your_api_key_here
   ```
4. Run the backend server:
   ```sh
   go run main.go
   ```

## API Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/live-flights` | Get real-time flight data with search & pagination |

### Example API Call
```sh
curl "http://localhost:8080/api/v1/live-flights?flightNumber=AA123&page=1&pageSize=10"
```

# Final Notes
- Ensure you have a valid **AviationStack API key**.
- The backend must be running before the frontend can fetch data.
- Future enhancements can include **websocket-based live updates**.

🚀 **Enjoy real-time flight tracking!**