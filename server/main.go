package main

import (
	"fmt"

	clientproxy "github.com/MatheusAlvesAlmeida/myMiddleware/distribution/client_proxy"
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/invoker"
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/shared"
	"github.com/MatheusAlvesAlmeida/myMiddleware/services/naming_service/proxy"
)

func main() {
	fmt.Println("Server running!")
	namingProxy := proxy.NamingProxy{}

	port := shared.FindNextAvailablePort()

	clientProxyPercentageCalculator := clientproxy.NewClientProxyPercentageCalculator("localhost", port, 1)

	namingProxy.Register("PercentageCalculator", clientProxyPercentageCalculator)

	invoker := invoker.Invoker{Port: port}

	invoker.Invoke()
}
