package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"runtime/debug"

	"github.com/otiai10/gosseract/v2"
)

// OCRResponse represents structured key-value output
type OCRResponse struct {
	// Fields map[string]string `json:"fields"`
	Text string `json:"text"`
}

func OCRHandler(w http.ResponseWriter, r *http.Request) {

	// 1️⃣ Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// --- Handle preflight OPTIONS request ---
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Only allow POST for actual OCR
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		debug.PrintStack()
		http.Error(w, "Failed to read image", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Ensure images folder exists
	err = os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		debug.PrintStack()
		http.Error(w, "Unable to create uploads folder", 500)
		return
	}

	uploadPath := "./uploads/" + header.Filename

	dst, err := os.Create(uploadPath)
	if err != nil {
		// Print a stack trace
		debug.PrintStack()
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		debug.PrintStack()
		http.Error(w, "Failed to write file contents: "+err.Error(), http.StatusInternalServerError)
		return
	}

	client := gosseract.NewClient()
	defer client.Close()

	client.SetImage(uploadPath)

	text, err := client.Text()
	if err != nil {
		debug.PrintStack()
		http.Error(w, "OCR processing failed", http.StatusInternalServerError)
		return
	}

	// Extract key-value pairs from cleaned text
	// fields := utils.ExtractKeyValues(text)

	response := OCRResponse{
		Text: text,
	}

	fmt.Print("test")
	fmt.Print(response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(OCRResponse{Text: text})

	// json.NewEncoder(w).Encode(response)
}
