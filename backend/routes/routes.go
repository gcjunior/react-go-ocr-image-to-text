package routes

import (
	"net/http"

	"ocr-api/handlers"
)

func SetupRoutes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/ocr", handlers.OCRHandler)

	return mux
}
