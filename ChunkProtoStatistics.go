package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto" // Importing the protobuf library
	"mygoapp/protoF" // Importing your protobuf-generated types
)

// Helper function to convert bytes to KB
func bytesToKB(sizeInBytes int) float64 {
	return float64(sizeInBytes) / 1024.0
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	// Read the incoming request body
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body_MST", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Received raw data size: %.2f KB\n", bytesToKB(len(reqData)))
	// Unmarshal the data into the Chunk message
	var req protoF.Chunk
	if err := proto.Unmarshal(reqData, &req); err != nil {
		http.Error(w, "Failed to unmarshal request_MST", http.StatusBadRequest)
		return
	}

	// Print the received data (optional, depending on your needs)
	// fmt.Println("Received Data:", string(req.Data))

	// Calculate and print the unmarshaled size
	unmarshaledSize := bytesToKB(len(reqData))
	fmt.Printf("Unmarshaled size: %.52 KB\n", unmarshaledSize)


	// Print the unmarshaled data for inspection (optional)
	// fmt.Printf("Unmarshaled Chunk: %+v\n", req)
fmt.Print(len(req.Data) )

	// Marshal the data again
	marshaledData, err := proto.Marshal(&req)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
    // Log the marshaled data to verify it
    // log.Println("Marshaled data:", marshaledData)
	// Calculate and print the marshaled size
	marshaledSize := bytesToKB(len(marshaledData))
	fmt.Printf("\nMarshaled size: %.2f KB\n", marshaledSize)
	fmt.Printf("________________________________________________________\n")

	// Send the marshaled data back as a response
	w.Header().Set("Content-Type", "application/protobuf")
	if _, err := w.Write(marshaledData); err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/chunk", greetHandler)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
