package main

import (
	"fmt"

	"github.com/MatheusAlvesAlmeida/myMiddleware/infrastructure/crh"
)

func main() {
	crh := crh.ClientRequestHandler{
		ServerHost: "localhost",
		ServerPort: "8080",
		Conn:       nil,
	}

	protocol := "tcp"

	message := []byte("Hello, Server!")
	response, err := crh.SendReceive(message, protocol)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Server Response:", string(response))
	}
}
