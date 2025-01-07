package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Handler function to handle POST requests
func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	// Read the byte data from the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Convert the byte data to string
	receivedString := string(body)

	// Print the received string
	fmt.Println("Received (converted to string):", receivedString)

	// Send a response back to the client
	w.Write([]byte("Message received successfully"))
}

func main() {
	// Define the route and associate it with the handler function
	http.HandleFunc("/receive", handlePostRequest)

	// Start the HTTP server on port 8080
	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
