package main

import (
	"fmt"
	"io"
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

// handleConnection reads the bytes from the connection, converts to string, and prints it
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read all data from the connection (instead of using Scanner)
	buffer := make([]byte, 1024) // Adjust size as needed
	n, err := conn.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Println("Error reading from connection:", err)
		return
	}

	// Convert the received byte data to string
	receivedString := string(buffer[:n])
	fmt.Println("Received (converted to string):", receivedString)

	// Optionally, you can send a response back to the client
	conn.Write([]byte("Message received\n"))
}
