/*****************************************************************************
 * client.go
 * Name: Sammy Lamothe and Jalen Jones
 * NetId: sl2938, jj9832
 *****************************************************************************/

package main

import (
	"io"
	"log"
	"net"
	"os"
)

const SEND_BUFFER_SIZE = 2048

// client opens a socket and sends message from stdin
func client(server_ip string, server_port string) {
	// Connect to the server
	conn, err := net.Dial("tcp", server_ip+":"+server_port)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create a buffer to read from stdin
	buf := make([]byte, SEND_BUFFER_SIZE)

	for {
		// Read data from stdin
		n, err := os.Stdin.Read(buf)
		if err != nil {
			// Handle EOF gracefully
			if err == io.EOF {
				break
			}
			log.Fatalf("Failed to read from stdin: %v", err)
		}

		// Send the data to the server in chunks
		totalSent := 0
		for totalSent < n {
			sent, err := conn.Write(buf[totalSent:n])
			if err != nil {
				log.Fatalf("Failed to send data: %v", err)
			}
			totalSent += sent
		}
	}
}

// Main parses command-line arguments and calls client function
func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: ./client [server IP] [server port] < [message file]")
	}
	server_ip := os.Args[1]
	server_port := os.Args[2]
	client(server_ip, server_port)
}
