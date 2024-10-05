/*****************************************************************************
 * client.go
 * Name: Sammy Lamothe and Jalen Jones
 * NetId: sl2938, jj9832
 *****************************************************************************/

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const SEND_BUFFER_SIZE = 2048

/* client()
 * Open socket and send message from stdin.
 */
func client(server_ip string, server_port string) {
	// Connect to the server
	conn, err := net.Dial("tcp", server_ip+":"+server_port)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create a buffer to read from stdin
	reader := bufio.NewReader(os.Stdin)

	// Read input from stdin and send it to the server
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read from stdin: %v", err)
		}

		// Send message to the server
		_, err = fmt.Fprintf(conn, text)
		if err != nil {
			log.Fatalf("Failed to send data: %v", err)
		}

		// Break after sending a message
		break
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
