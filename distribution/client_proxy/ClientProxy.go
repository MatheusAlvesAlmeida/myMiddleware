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
	Requestor *requestor.Requestor // Add a field for the Requestor
}

type ClientProxyPercentageCalculator struct {
	Proxy PercentageProxy
}

func NewPercentageProxy(aor int) PercentageProxy {
	return PercentageProxy{
		ID:        aor,
		Host:      "localhost",
		Port:      8080,
		Requestor: &requestor.Requestor{},
	}
}

func NewClientProxyPercentageCalculator(host string, port int, id int) ClientProxyPercentageCalculator {
	typeName := reflect.TypeOf(ClientProxyPercentageCalculator{}).String()
	return ClientProxyPercentageCalculator{
		PercentageProxy{TypeName: typeName, Host: host, Port: port, ID: id}}
}

func (proxy PercentageProxy) GetValueOf(percentage int, totalValue int) float64 {
	params := make([]interface{}, 2)
	params[0] = percentage
	params[1] = totalValue

	request := shared.Request{Op: "GetValueOf", Params: params}
	invoker := shared.Invocation{Host: proxy.Host, Port: proxy.Port, Request: request}

	response := proxy.Requestor.Invoke(invoker).([]interface{})

	return response[0].(float64)
}

func (proxy PercentageProxy) GetPercentageOf(partialValue int, totalValue int) float64 {
	params := make([]interface{}, 2)
	params[0] = partialValue
	params[1] = totalValue

	request := shared.Request{Op: "GetPercentageOf", Params: params}
	invoker := shared.Invocation{Host: proxy.Host, Port: proxy.Port, Request: request}

	response := proxy.Requestor.Invoke(invoker).([]interface{})

	return response[0].(float64)
}
