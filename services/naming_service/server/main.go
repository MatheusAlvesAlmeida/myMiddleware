package main

import (
	"github.com/MatheusAlvesAlmeida/myMiddleware/services/naming_service/invoker"
)

func main() {
	namingInvoker := invoker.NamingInvoker{}
	namingInvoker.Invoke()
}
