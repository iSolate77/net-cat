package server

import (
	"log"
	"net"
)

// HandleConnections manages incoming client connections.
func HandleConnections(listener *net.TCPListener) {
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		go handleClient(conn)
	}
}

// handleClient manages a single client connection.
func handleClient(conn *net.TCPConn) {
	defer conn.Close()
	log.Printf("Client connected: %s", conn.RemoteAddr().String())

	// Implement the logic for handling client communication here

	log.Printf("Client disconnected: %s", conn.RemoteAddr().String())
}
