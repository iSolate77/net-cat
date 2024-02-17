package main

import (
	// "bufio"
	"fmt"
	// "net"
	"github.com/iSolate77/net-cat/internal/client"
	"os"
)

func main() {
	serverAddress := "localhost:8989"
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./client $serverAddress")
		os.Exit(1)
	} else if len(os.Args) == 2 {
		serverAddress = os.Args[1]
	}

	client.ConnectToServer(serverAddress)
}
