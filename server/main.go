package main

import (
	"github.com/MatheusAlvesAlmeida/myMiddleware/infrastructure/srh"
)

func main() {
	server := srh.ServerRequestHandler{
		ServerHost: "localhost",
		ServerPort: "8080",
		Protocol:   "tcp",
	}

	err := server.Server()
	if err != nil {
		panic(err)
	}
}
