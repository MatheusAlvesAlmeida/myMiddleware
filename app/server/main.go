package main

import (
	clientproxy "github.com/MatheusAlvesAlmeida/myMiddleware/distribution/client_proxy"
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/invoker"
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/shared"
	"github.com/MatheusAlvesAlmeida/myMiddleware/services/naming_service/proxy"
)

func main() {
	namingProxy := proxy.NamingProxy{}

	port := shared.FindNextAvailablePort()

	clientProxyPercentageCalculator := clientproxy.NewPercentageProxyCalculator("localhost", port, 1)

	namingProxy.Register("PercentageCalculator", clientProxyPercentageCalculator)

	myInvoker := invoker.Invoker{Port: port}
	myInvoker.Invoke()
}
