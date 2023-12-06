package repository

import (
	"reflect"

	clientproxy "github.com/MatheusAlvesAlmeida/myMiddleware/distribution/client_proxy"
)

func CheckRepository(proxy clientproxy.PercentageProxy) interface{} {
	var clientProxy interface{}

	switch proxy.TypeName {
	case reflect.TypeOf(clientproxy.ClientProxyPercentageCalculator{}).String():
		newProxy := clientproxy.NewPercentageProxyCalculator(proxy.Host, proxy.Port, proxy.ID)
		newProxy.Proxy.Host = proxy.Host
		newProxy.Proxy.Port = proxy.Port
		newProxy.Proxy.TypeName = proxy.TypeName

		clientProxy = newProxy
	}

	return clientProxy
}
