package clientproxy

import (
	"reflect"

	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/requestor"
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/shared"
)

type PercentageProxy struct {
	ID        int
	Host      string
	Port      int
	TypeName  string
	Requestor *requestor.Requestor
}

type ClientProxyPercentageCalculator struct {
	Proxy PercentageProxy
}

func NewPercentageProxy(host string, port int, id int) PercentageProxy {
	return PercentageProxy{
		ID:        id,
		Host:      host,
		Port:      port,
		Requestor: &requestor.Requestor{},
	}
}

func NewPercentageProxyCalculator(host string, port int, id int) ClientProxyPercentageCalculator {
	typeName := reflect.TypeOf(ClientProxyPercentageCalculator{}).String()
	return ClientProxyPercentageCalculator{
		PercentageProxy{TypeName: typeName, Host: host, Port: port, ID: id, Requestor: &requestor.Requestor{}},
	}
}

func (proxy ClientProxyPercentageCalculator) GetValueOf(percentage int, totalValue int) float64 {
	params := make([]interface{}, 2)
	params[0] = percentage
	params[1] = totalValue

	request := shared.Request{Op: "GetValueOf", Params: params}
	invoker := shared.Invocation{Host: proxy.Proxy.Host, Port: proxy.Proxy.Port, Request: request}

	response := proxy.Proxy.Requestor.Invoke(invoker).([]interface{})

	return response[0].(float64)
}

func (proxy ClientProxyPercentageCalculator) GetPercentageOf(partialValue int, totalValue int) float64 {
	params := make([]interface{}, 2)
	params[0] = partialValue
	params[1] = totalValue

	request := shared.Request{Op: "GetPercentageOf", Params: params}
	invoker := shared.Invocation{Host: proxy.Proxy.Host, Port: proxy.Proxy.Port, Request: request}

	response := proxy.Proxy.Requestor.Invoke(invoker).([]interface{})

	return response[0].(float64)
}
