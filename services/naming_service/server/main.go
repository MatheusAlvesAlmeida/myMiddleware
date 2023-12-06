package main

import (
	"fmt"

	"github.com/MatheusAlvesAlmeida/myMiddleware/services/naming_service/invoker"
)

func main() {

	fmt.Println("Naming service is running!")

	namingInvoker := invoker.NamingInvoker{}
	namingInvoker.Invoke()
}
