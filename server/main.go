package main

import (
	"fmt"

	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/invoker"
)

func main() {
	fmt.Println("Server running!")

	invoker := invoker.NewInvoker()
	invoker.Invoke()
}
