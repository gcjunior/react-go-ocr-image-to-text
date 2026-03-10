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

	log.Printf("Front End URL %s\n", frontendURL)

	router := routes.SetupRoutes()

	handler := middleware.Logger(
		middleware.CORS(router, frontendURL),
	)

	log.Printf("Server running on port %s\n", port)

	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal(err)
	}
}
