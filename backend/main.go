package main

import (
	"log"
	"net/http"

	"ocr-api/middleware"
	"ocr-api/routes"
)

func main() {

	router := routes.SetupRoutes()

	handler := middleware.Logger(
		middleware.CORS(router),
	)

	log.Println("🚀 Server running on :8080")

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}
