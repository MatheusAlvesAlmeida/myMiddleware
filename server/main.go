package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MatheusAlvesAlmeida/myMiddleware/infrastructure/srh"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: main <protocol>")
		return
	}

	protocol := args[0]

	var handlerTCP *srh.ServerRequestHandlerTCP
	var handlerUDP *srh.ServerRequestHandlerUDP

	switch protocol {
	case "tcp":
		handlerTCP = &srh.ServerRequestHandlerTCP{}
	case "udp":
		handlerUDP = &srh.ServerRequestHandlerUDP{}
	default:
		fmt.Println("Unsupported protocol")
		return
	}

	if handlerTCP != nil {
		server := handlerTCP.Server("127.0.0.1", "8080")
		response, err := server.ReceiveMessage()

		if err != nil {
			fmt.Printf("Error receiving TCP message: %v\n", err)
		} else {
			fmt.Printf("Debug info - Received TCP message: %s\n", string(response))
			response := ping(string(response))
			server.SendMessage(response)
		}
	} else {
		server := handlerUDP.Server("127.0.0.1", "8080")
		response, clientAddr, err := server.ReceiveMessage()

		if err != nil {
			fmt.Printf("Error receiving UDP message: %v\n", err)
		} else {
			fmt.Printf("Debug info - Received UDP message: %s\n", string(response))
			response := ping(string(response))
			fmt.Println("Debug info - Response in server", string(response))
			server.SendMessage(clientAddr, response)
		}
	}
}

func ping(url string) []byte {
	fmt.Println("Pinging", url, "...")

	resp, err := http.Get(url)
	if err != nil {
		return []byte("Failed to ping the website: " + err.Error())
	}
	defer resp.Body.Close()

	var response string

	if resp.StatusCode == http.StatusOK {
		response = "Website " + url + " is up!"
	} else {
		response = "Website " + url + " is down!"
	}

	return []byte(response)
}
