# 🛫 Airline Tracking Service

## 🚀 Overview
This application provides real-time flight tracking using live data from the AviationStack API. It features a **React frontend** and a **Golang backend** with Redis caching, WebSockets for real-time updates, and interactive flight mapping using **Leaflet.js**.

---

## 🌟 Features

### ✅ **Real-Time Flight Tracking**
- Fetches live flight data with latitude & longitude.
- Displays real-time flight positions on an interactive map.
- Uses Leaflet.js for visualization.

### ✅ **Redis Caching**
- Speeds up API responses by caching recent flight data.
- Reduces redundant external API calls.

### ✅ **WebSockets for Live Updates**
- Automatically updates flights in real-time.
- No need to manually refresh the page.

### ✅ **Flight Search & Filtering**
- Search by **flight number, airline, departure, or arrival**.
- Supports **pagination** for better data handling.

### ✅ **Auto-Suggestions for Airport Codes**
- Fetches airport codes dynamically as users type.
- Uses an external API for auto-suggestions.

### ✅ **Recent Search History**
- Stores past searches in **local storage**.
- Allows users to quickly repeat previous searches.

### ✅ **Dark Mode Support** 🌙
- Toggle between **light and dark mode** for better user experience.

### ✅ **Fully Containerized (Docker & Docker Compose)**
- Runs both backend & frontend using Docker.
- Includes **Redis** for caching.

---

## 📦 Installation & Setup

### 🔹 **1. Clone the Repository**
```bash
git clone https://github.com/your-repo/airline-tracking.git
cd airline-tracking
```

### 🔹 **2. Set Up Environment Variables**
Create a `.env` file in the **backend** directory:
```plaintext
AVIATIONSTACK_API_KEY=your_api_key_here
REDIS_ADDR=localhost:6379
```

### 🔹 **3. Run with Docker Compose**
```bash
docker-compose up --build
```

---

## 🖥️ Frontend (React)
- Accessible at: **http://localhost:3000**
- Uses **React + Tailwind CSS** for styling.
- Implements **Dark Mode**, **Search History**, and **Live Flight Mapping**.

## ⚙️ Backend (Golang)
- Runs on **http://localhost:8080**.
- Uses **Redis for caching** & **WebSockets for live updates**.

---

## 📌 API Endpoints
### 🔹 **1. Fetch Live Flights**
```http
GET /api/v1/live-flights?page=1&pageSize=10
```
**Query Parameters:**
- `flightNumber`: Filter by flight number
- `airline`: Filter by airline name
- `departure`: Filter by departure airport
- `arrival`: Filter by arrival airport

### 🔹 **2. Fetch Flights with Mapping Data**
```http
GET /api/v1/live-flights?latitude&longitude
```
- Returns **real-time flight positions** for mapping.

---

## 🌍 Future Enhancements
- ✅ **User Authentication** (Login & Save Favorite Flights)
- ✅ **Weather Integration** for Flight Routes
- ✅ **Email & Push Notifications for Delayed Flights**

---

## 🤝 Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

---

## 📜 License
This project is licensed under the **MIT License**.