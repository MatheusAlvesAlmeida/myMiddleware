package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/MatheusAlvesAlmeida/myMiddleware/infrastructure/srh"
)

func main() {

	var handlerTCP *srh.ServerRequestHandlerTCP

	handlerTCP = &srh.ServerRequestHandlerTCP{}

	server := handlerTCP.Server("127.0.0.1", "8080")
	fmt.Println("Server is listening on port 8080")
	for {
		response, err := server.ReceiveMessage()

		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected")
			} else {
				fmt.Printf("Error receiving TCP message: %v\n", err)
			}
			break
		}

		fmt.Printf("Debug info - Received TCP message: %s\n", string(response))
		response = ping(string(response))
		server.SendMessage(response)
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
