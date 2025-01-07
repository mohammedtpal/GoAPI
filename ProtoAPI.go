package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/golang/protobuf/proto"  // Importing the Go protobuf library
	"io/ioutil"
	"mygoapp/protoF" // Corrected import for your generated protobuf code
)

func handleGreeting(w http.ResponseWriter, r *http.Request) {
	// Read the incoming Protobuf data
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	// Unmarshal the Protobuf data into GreetingRequest
	var req protoF.GreetingRequest // Use the correct package and type
	if err := proto.Unmarshal(data, &req); err != nil {
		http.Error(w, "Failed to unmarshal Protobuf", http.StatusBadRequest)
		return
	}
	

	// Create a response message
	res := &protoF.GreetingResponse{ // Use the correct package and type
		Message: "Hello, " + req.Name,
	}

	// Marshal the response back to Protobuf
	resData, err := proto.Marshal(res)
	if err != nil {
		http.Error(w, "Failed to marshal Protobuf response", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to "application/protobuf"
	w.Header().Set("Content-Type", "application/protobuf")
	w.Write(resData)
}

func main() {
	http.HandleFunc("/greet", handleGreeting)

	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
