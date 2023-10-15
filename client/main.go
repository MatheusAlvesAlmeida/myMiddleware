package main

import (
	"fmt"

	"github.com/AP/myMiddleware/infrastructure/crh"
)

func main() {
	crh := crh.ClientRequestHandler{
		ServerHost: "localhost",
		ServerPort: "8080",
		nil,
	}

	protocol := "tcp"

	message := []byte("Hello, Server!")
	response := crh.SendReceive(message, protocol)

	fmt.Println("Server Response:", string(response))
}
