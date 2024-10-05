/*****************************************************************************
 * server.go
 * Name:
 * NetId:
 *****************************************************************************/

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const SERVER_PORT = "8000"

/* server()
 * Open socket, accept client connections, and process messages.
 */
func server(server_ip string, server_port string) {
	// Combine IP and port into a single address string
	address := fmt.Sprintf("%s:%s", server_ip, server_port)

	// Create a TCP listener on the specified address
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer ln.Close()

	fmt.Printf("Server listening on %s\n", address)

	// Infinite loop to accept and handle client connections
	for {
		// Wait for a connection attempt from a client
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue // Try to accept another connection
		}

		// Handle client connection in a new goroutine (optional for concurrency)
		go handleClient(conn)
	}
}

/* handleClient()
 * Reads data from client, processes it, and sends back a response.
 */
func handleClient(conn net.Conn) {
	defer conn.Close() // Ensure connection is closed after processing

	fmt.Printf("Connected to %s\n", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)
	for {
		// Read client message
		message, err := reader.ReadString('\n')
		if err != nil {
			// Check if the error is EOF (client closed connection)
			if err.Error() == "EOF" {
				fmt.Println("Client disconnected")
				return
			}
			log.Printf("Error reading message: %v", err)
			return
		}

		// Print received message from client
		fmt.Printf("Message Received: %s", message)

		// Process the message (convert to uppercase)
		newMessage := strings.ToUpper(message)

		// Send processed message back to client
		_, err = conn.Write([]byte(newMessage + "\n"))
		if err != nil {
			log.Printf("Failed to send response: %v", err)
			return
		}
	}
}

// Main function to parse command-line arguments and start the server
func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: ./server [server IP]")
	}
	server_ip := os.Args[1]
	server(server_ip, SERVER_PORT)
}
