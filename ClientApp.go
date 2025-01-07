package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto" // Importing the protobuf library
	"mygoapp/protoF" // Importing your protobuf-generated types
)

func main() {
	// Create a GreetingRequest message
	req := &protoF.GreetingRequest{
		Name: "John Doe",
	}

	// Marshal the request into Protobuf format
	data, err := proto.Marshal(req)
	if err != nil {
		log.Fatalf("Failed to marshal request: %v", err)
	}

	// Send the request to the server
	resp, err := http.Post("http://localhost:8080/greet", "application/protobuf", bytes.NewReader(data))
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: received non-OK status code %d", resp.StatusCode)
	}

	// Unmarshal the response
	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var res protoF.GreetingResponse
	if err := proto.Unmarshal(resData, &res); err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Print the response message
	fmt.Println("Server Response:", res.Message)
}
