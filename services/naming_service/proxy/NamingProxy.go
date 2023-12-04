package proxy

import (
	clientproxy "github.com/MatheusAlvesAlmeida/myMiddleware/distribution/client_proxy"
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/requestor"
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/shared"
	"github.com/MatheusAlvesAlmeida/myMiddleware/repository"
)

type NamingProxy struct{}

func (NamingProxy) Register(name string, proxy interface{}) bool {
	params := make([]interface{}, 2)
	params[0] = name
	params[1] = proxy

	namingProxy := clientproxy.PercentageProxy{ID: 0, Host: "", Port: shared.NAMING_PORT, TypeName: "NamingProxy"}
	request := shared.Request{Op: "Register", Params: params}
	inv := shared.Invocation{Host: namingProxy.Host, Port: namingProxy.Port, Request: request}

	requestor := requestor.Requestor{}
	invoker := requestor.Invoke(inv)

	response, _ := invoker.([]interface{})

	result, _ := response[0].(bool)

	return result
}

func (NamingProxy) Lookup(name string) interface{} {
	params := make([]interface{}, 1)
	params[0] = name

	namingProxy := clientproxy.PercentageProxy{ID: 0, Host: "", Port: shared.NAMING_PORT, TypeName: "NamingProxy"}
	request := shared.Request{Op: "Lookup", Params: params}
	inv := shared.Invocation{Host: namingProxy.Host, Port: namingProxy.Port, Request: request}

	requestor := requestor.Requestor{}
	invoker := requestor.Invoke(inv)

	response, _ := invoker.([]interface{})

	result, _ := response[0].(map[string]interface{})

	proxyTemp := result
	clientProxyTemp := clientproxy.PercentageProxy{TypeName: proxyTemp["TypeName"].(string), Host: proxyTemp["Host"].(string), Port: int(proxyTemp["Port"].(float64))}
	clientProxy := repository.CheckRepository(clientProxyTemp)

	return clientProxy
}
