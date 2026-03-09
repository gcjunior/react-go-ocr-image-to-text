package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"runtime/debug"

	"github.com/otiai10/gosseract/v2"
)

type OCRResponse struct {
	Text string `json:"text"`
}

func OCRHandler(w http.ResponseWriter, r *http.Request) {

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

	response := OCRResponse{
		Text: text,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
