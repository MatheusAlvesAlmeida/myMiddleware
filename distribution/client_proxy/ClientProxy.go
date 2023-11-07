package clientproxy

import (
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/requestor"
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/shared"
)

type PercentageProxy struct {
	Host string
	Port int
}

func NewPercentageProxy(host string, port int) PercentageProxy {
	return PercentageProxy{host, port}
}

func (proxy PercentageProxy) GetValueOf(percentage int, totalValue int) float64 {
	requestor := requestor.Requestor{}

	params := make([]interface{}, 2)
	params[0] = percentage
	params[1] = totalValue

	request := shared.Request{Op: "GetValueOf", Params: params}
	invoker := shared.Invocation{Host: proxy.Host, Port: proxy.Port, Request: request}

	response := requestor.Invoke(invoker).([]interface{})

	return response[0].(float64)
}

func (proxy PercentageProxy) GetPercentageOf(partialValue int, totalValue int) float64 {
	requestor := requestor.Requestor{}

	params := make([]interface{}, 2)
	params[0] = partialValue
	params[1] = totalValue

	request := shared.Request{Op: "GetPercentageOf", Params: params}
	invoker := shared.Invocation{Host: proxy.Host, Port: proxy.Port, Request: request}

	response := requestor.Invoke(invoker).([]interface{})

	return response[0].(float64)
}
