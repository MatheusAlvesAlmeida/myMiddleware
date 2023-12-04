package main

import (
	"fmt"

	"github.com/MatheusAlvesAlmeida/myMiddleware/services/naming_service/invoker"
)

func main() {

	fmt.Println("Naming servidor running!!")

	// control loop passed to invoker
	namingInvoker := invoker.NamingInvoker{}
	namingInvoker.Invoke()
}
