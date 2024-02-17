package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// ConnectToServer handles the client's connection to the server.
func ConnectToServer(serverAddress string) {
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to server. Enter messages to send.")

	// Start a goroutine for reading server responses
	go readFromServer(conn)

	// Read from stdin and send to the server
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		_, err := conn.Write([]byte(text + "\n"))
		if err != nil {
			fmt.Printf("Error sending message: %v\n", err)
			break
		}
	}
}

// readFromServer reads messages from the server and prints them
func readFromServer(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println(message)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading from server: %v\n", err)
	}
}
