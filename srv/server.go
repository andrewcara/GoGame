// server.go (Backend - runs on :8080)
package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	setupAPI(mux)

	// Configure CORS to allow frontend requests
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5501"}, // Frontend URL
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"*"},
	})
	handler := c.Handler(mux)
	log.Println("Backend server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func setupAPI(mux *http.ServeMux) {
	// WebSocket handler
	manager := NewManager()
	mux.HandleFunc("/ws", manager.serveWS)

	// Any other API endpoints
	mux.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running"))
	})
}
