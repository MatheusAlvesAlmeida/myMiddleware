package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MatheusAlvesAlmeida/myMiddleware/infrastructure/crh"
)

func main() {

	crh := crh.ClientRequestHandlerTCP{
		ServerAddress: "localhost:8080",
		Conn:          nil,
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a message (type 'end' to quit): ")
		scanner.Scan()
		input := scanner.Text()

		if input == "end" {
			break
		}

		message := []byte(input)
		response, err := crh.SendReceive(message)
		if err != nil {
			fmt.Printf("Error sending message: %s\n", err)
			return
		}
		fmt.Printf("Response from server: %s\n", string(response))
	}

}
