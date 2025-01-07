package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto" // Importing the protobuf library
	"mygoapp/protoF" // Importing your protobuf-generated types
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	// Read the incoming request body
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body_MST", http.StatusInternalServerError)
		return
	}

	// Unmarshal the data into the Chunk message
	var req protoF.Chunk
	if err := proto.Unmarshal(reqData, &req); err != nil {
		http.Error(w, "Failed to unmarshal request_MST", http.StatusBadRequest)
		return
	}

	// You can log the received data or do any processing here
	fmt.Println("Received Data:", string(req.Data))

	// Send the same data back as a response
	w.Header().Set("Content-Type", "application/protobuf")
	if err := proto.MarshalText(w, &req); err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/chunk", greetHandler)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
