package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	clientproxy "github.com/MatheusAlvesAlmeida/myMiddleware/distribution/client_proxy"
	"github.com/MatheusAlvesAlmeida/myMiddleware/services/naming_service/proxy"
)

const MyAOR = 0

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	namingProxy := proxy.NamingProxy{}

	proxy := namingProxy.Lookup("PercentageCalculator").(clientproxy.PercentageProxy)

	for {
		fmt.Println("Enter the operation you want to perform (type 'end' to quit): ")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "GetValueOf":
			fmt.Println("Enter the percentage and the total value: ")
			scanner.Scan()
			percentage := scanner.Text()
			scanner.Scan()
			totalValue := scanner.Text()

			percentageInt, _ := strconv.Atoi(percentage)
			totalValueInt, _ := strconv.Atoi(totalValue)

			fmt.Println("Response from server: ", proxy.GetValueOf(percentageInt, totalValueInt))
		case "GetPercentageOf":
			fmt.Println("Enter the partial value and the total value: ")
			scanner.Scan()
			partialValue := scanner.Text()
			scanner.Scan()
			totalValue := scanner.Text()

			partialValueInt, _ := strconv.Atoi(partialValue)
			totalValueInt, _ := strconv.Atoi(totalValue)

			fmt.Println("Response from server: ", proxy.GetPercentageOf(partialValueInt, totalValueInt))
		case "end":
			os.Exit(0)
		default:
			fmt.Println("Invalid operation")
		}
	}
}
