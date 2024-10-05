/*****************************************************************************
 * server.go
 * Name:
 * NetId:
 *****************************************************************************/

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const RECV_BUFFER_SIZE = 2048

/* TODO: server()
 * Open socket and wait for client to connect
 * Print received message to stdout
 */
func server(server_port string) {
	address := fmt.Sprintf("127.0.0.1:%s", server_port)

	// Start to listen for incoming connections
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error starting the server on port %s: %v", server_port, err)
	}
	fmt.Printf("Server is listening on %s\n", address)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting client connection: %v", err)
			continue // Move on to the next connection attempt
		}
		// Handle the connected client in a goroutine for concurrency
		go handleConnection(conn)
	}
}

/* TODO: handleConnection()
 * Handle an individual client connection
 */
func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		// Create a buffer and read the message
		message, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Client at %s disconnected\n", conn.RemoteAddr().String())
				break
			}
			log.Printf("Error reading from client: %v", err)
			break
		}
		fmt.Print("Message received:", message)
	}
}

// Main parses command-line arguments and calls server function
func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: ./server [server port]")
	}
	server_port := os.Args[1]
	server(server_port)
}
