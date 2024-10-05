/*****************************************************************************
 * server.go
 * Name:
 * NetId:
 *****************************************************************************/

package main

import (
	"fmt"
	"io"
	"log"
	"net"
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
		// handle error
		log.Fatalf("Error starting the server on port %s: %v", server_port, err)
	}

	fmt.Printf("Server is listening on %s\n", address)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting client connection: %v", err)
			continue // Move on to the next connection attempt
		}
		// Handle the connected client in a dedicated function
		handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	// Make sure the connection is closed once we're done
	defer conn.Close()
	// Create a buffer to store the received data
	buf := make([]byte, RECV_BUFFER_SIZE)

	// Keep reading data from the client until they disconnect or there's an error
	for {
		n, err := conn.Read(buf)
		if err != nil {
			// Handle EOF gracefully
			if err == io.EOF {
				fmt.Printf("Client at %s disconnected\n", conn.RemoteAddr().String())
				break
			}
			log.Printf("error reading from client: %v", err)

		}
		buf = buf[:n]
		fmt.Print(string(buf[:n]))
	}
}
