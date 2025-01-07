package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Set up the TCP server to listen on port 8080
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listen.Close()

	fmt.Println("Server listening on port 8080...")

	// Accept incoming connections
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the incoming connection in a new goroutine
		go handleConnection(conn)
	}
}

// handleConnection reads the string from the connection and prints it
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read data from the connection
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		receivedString := scanner.Text()
		fmt.Println("Received:", receivedString)

		// Optionally, you can send a response back to the client
		conn.Write([]byte("Message received\n"))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from connection:", err)
	}
}
