package clientproxy

import (
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/requestor"
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/shared"
)

type PercentageProxy struct {
	AOR       string
	Host      string
	Port      int
	Requestor *requestor.Requestor // Add a field for the Requestor
}

func NewPercentageProxy(aor string) PercentageProxy {
	return PercentageProxy{
		AOR:       aor,
		Host:      "localhost",
		Port:      8080,
		Requestor: &requestor.Requestor{},
	}
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
