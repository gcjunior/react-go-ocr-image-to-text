package main

import (
	"log"
	"net/http"
	"os"

	"ocr-api/middleware"
	"ocr-api/routes"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000"
	}

	router := routes.SetupRoutes(frontendURL)

	handler := middleware.Logger(
		middleware.CORS(router),
	)

	log.Println("🚀 Server running on :8080")

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}
