package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Helper function to convert bytes to KB
func bytesToKB(sizeInBytes int) float64 {
	return float64(sizeInBytes) / 1024
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	// Read the incoming request body (which is expected to be base64 encoded)
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// Remove any extra whitespace or newlines and base64 decode the data
	decodedData, err := base64.StdEncoding.DecodeString(strings.TrimSpace(string(reqData)))
	if err != nil {
		http.Error(w, "Failed to decode base64", http.StatusBadRequest)
		return
	}
	Base64Size := bytesToKB(len(reqData))
	fmt.Printf("Size of Base64Size bytes: %.2f KB\n", Base64Size)

	// Print the size of the decoded raw bytes
	rawSize := bytesToKB(len(decodedData))
	fmt.Printf("Size of raw decoded bytes: %.2f KB\n", rawSize)

	// For demonstration, printing the decoded bytes as a string (optional)
	// Be careful as this might print non-printable characters depending on the content
	// fmt.Println("Decoded Raw Data (String representation):", string(decodedData))

	// Respond back with the decoded raw data (or handle further processing as required)
	w.Header().Set("Content-Type", "application/octet-stream")
	if _, err := w.Write(decodedData); err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/chunk", greetHandler)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
