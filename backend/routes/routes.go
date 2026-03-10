package routes

import (
	"net/http"

	"ocr-api/handlers"
)

func SetupRoutes(frontendURL string) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/ocr", handlers.OCRHandler(frontendURL))

	return mux
}
