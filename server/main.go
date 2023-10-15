package main

import (
	"fmt"

	"github.com/MatheusAlvesAlmeida/myMiddleware/infrastructure/srh"
)

func main() {
	//protocol := "tcp"
	protocol := "udp"

	var handler srh.ServerRequestHandler

	switch protocol {
	case "tcp":
		handler = &srh.TCPHandler{
			SRH: &srh.ServerRequestHandlerTCP{Protocol: "tcp"},
		}
	case "udp":
		handler = &srh.UDPHandler{
			SRH: &srh.ServerRequestHandlerUDP{Protocol: "udp"},
		}
	default:
		fmt.Println("Unsupported protocol")
		return
	}

	if err := handler.Server("127.0.0.1", "8080"); err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
	}
}
