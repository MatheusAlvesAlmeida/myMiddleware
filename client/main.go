package main

import (
	"fmt"
	"os"

	"github.com/MatheusAlvesAlmeida/myMiddleware/infrastructure/crh"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Usage: main <protocol> <message>")
		return
	}

	protocol := args[0]
	message := []byte(args[1])

	switch protocol {
	case "tcp":
		crh := crh.ClientRequestHandlerTCP{
			ServerAddress: "localhost:8080",
			Conn:          nil,
		}
		response, err := crh.SendReceive(message)
		if err != nil {
			fmt.Printf("Error sending message: %s\n", err)
			return
		}
		fmt.Printf("Response from server: %s\n", string(response))
	case "udp":
		crh := crh.ClientRequestHandlerUDP{
			ServerAddress: "localhost:8080",
			Conn:          nil,
		}
		response, err := crh.SendReceive(message)
		if err != nil {
			fmt.Printf("Error sending message: %s\n", err)
			return
		}
		fmt.Printf("Response from server: %s\n", string(response))
	default:
		fmt.Println("Unsupported protocol")
		return
	}

}
