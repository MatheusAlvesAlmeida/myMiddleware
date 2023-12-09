package main

import (
	"fmt"

	clientproxy "github.com/MatheusAlvesAlmeida/myMiddleware/distribution/client_proxy"
	"github.com/MatheusAlvesAlmeida/myMiddleware/services/naming_service/proxy"
)

const MyAOR = 0

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	namingProxy := proxy.NamingProxy{}
	proxy := namingProxy.Lookup("PercentageCalculator").(clientproxy.ClientProxyPercentageCalculator)

	fmt.Println("Client running!")

	percentageValue, totalValue, partialValue := 22, 30, 22

	for i := 0; i < 100000; i++ {
		value, err := proxy.GetValueOf(percentageValue, totalValue)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("GetValueOf: ", value)
	}

	for i := 0; i < 100000; i++ {
		value, err := proxy.GetPercentageOf(partialValue, totalValue)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("GetPercentageOf: ", value)
	}
}
