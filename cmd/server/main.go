package main

import (
	"github.com/iSolate77/net-cat/internal/server"
	"log"
	"net"
	"os"
)

func main() {
	port := "8989" // Default port
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	serverAddress := "0.0.0.0:" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Error resolving TCP address: %v", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalf("Error listening on %s: %v", serverAddress, err)
	}
	defer listener.Close()
	log.Printf("Listening on %s", serverAddress)

	server.HandleConnections(listener)
}
