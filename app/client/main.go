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
	proxy := namingProxy.Lookup("PercentageCalculator").(clientproxy.ClientProxyPercentageCalculator)

	fmt.Println("Client running!")
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

			percentageInt, err := strconv.Atoi(percentage)
			if err != nil {
				fmt.Println("Invalid percentage input:", err)
				continue
			}

			totalValueInt, err := strconv.Atoi(totalValue)
			if err != nil {
				fmt.Println("Invalid total value input:", err)
				continue
			}

			value, err := proxy.GetValueOf(percentageInt, totalValueInt)
			if err != nil {
				fmt.Println("Error getting value from server:", err)
				continue
			}

			fmt.Println("Response from server: ", value)
		case "GetPercentageOf":
			fmt.Println("Enter the partial value and the total value: ")
			scanner.Scan()
			partialValue := scanner.Text()
			scanner.Scan()
			totalValue := scanner.Text()

			partialValueInt, err := strconv.Atoi(partialValue)
			if err != nil {
				fmt.Println("Invalid partial value input:", err)
				continue
			}

			totalValueInt, err := strconv.Atoi(totalValue)
			if err != nil {
				fmt.Println("Invalid total value input:", err)
				continue
			}

			value, err := proxy.GetPercentageOf(partialValueInt, totalValueInt)
			if err != nil {
				fmt.Println("Error getting percentage from server:", err)
				continue
			}

			fmt.Println("Response from server: ", value)
		case "end":
			os.Exit(0)
		default:
			fmt.Println("Invalid operation")
		}
		fmt.Printf("\n\n------------------------------------------------------\n")
	}
}
