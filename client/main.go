package main

import (
	"fmt"

	clientproxy "github.com/MatheusAlvesAlmeida/myMiddleware/distribution/client_proxy"
)

const MyAOR = "localhost:8080"

func main() {
	proxy := clientproxy.NewPercentageProxy(MyAOR)

	percentage := 20
	totalValue := 1000
	value := proxy.GetValueOf(percentage, totalValue)
	fmt.Printf("Value of %d%% of %d is %.2f\n", percentage, totalValue, value)

	partialValue := 300
	percentageValue := proxy.GetPercentageOf(partialValue, totalValue)
	fmt.Printf("%d is %.2f%% of %d\n", partialValue, percentageValue, totalValue)
}
