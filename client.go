/*****************************************************************************
 * client.go
 * Name: Sammy Lamothe and Jalen Jones
 * NetId: sl2938, jj9832
 *****************************************************************************/

package main

import (
	"log"
	"os"
)

const SEND_BUFFER_SIZE = 2048

// client opens a socket and sends message from stdin
func client(server_ip string, server_port string) {
	// Attempt to connect to the server using a raw connection
	conn, err := os.OpenFile("/dev/tcp/"+server_ip+"/"+server_port, os.O_RDWR, 0)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Buffer to hold chunks of data from stdin
	buf := make([]byte, SEND_BUFFER_SIZE)

	// Keep reading from stdin until there's nothing left (EOF)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			// If EOF is reached, stop the loop
			if err.Error() == "EOF" {
				break
			}
			// For any other error, terminate the program
			log.Fatalf("Failed to read from stdin: %v", err)
		}

		// Send the data to the server in chunks if necessary
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
