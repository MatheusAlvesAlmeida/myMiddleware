package invoker

import (
	"strconv"

	clientproxy "github.com/MatheusAlvesAlmeida/myMiddleware/distribution/client_proxy"
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/marshaller"
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/miop"
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/requestor"
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/shared"
	"github.com/MatheusAlvesAlmeida/myMiddleware/infrastructure/srh"
	naming "github.com/MatheusAlvesAlmeida/myMiddleware/services/naming_service"
)

type NamingInvoker struct{}

func (NamingInvoker) Invoke() {
	srh := srh.ServerRequestHandlerTCP{ServerHost: "localhost", ServerPort: strconv.Itoa(shared.NAMING_PORT)}
	marshaller := marshaller.Marshaller{}
	namingService := naming.NamingService{}
	miopPacketReply := miop.Packet{}
	responseParams := make([]interface{}, 1)

	for {
		messageReceived, err := srh.ReceiveMessage()
		if err != nil {
			continue
		}

		miopPacketRequest := marshaller.Unmarshall(messageReceived)
		operation := miopPacketRequest.Body.ReqHeader.Operation

		switch operation {
		case "Register":
			params := miopPacketRequest.Body.ReqBody.Body
			name := params[0].(string)
			proxy := params[1].(map[string]interface{})
			proxyTemp := proxy["Proxy"].(map[string]interface{})

			p2 := clientproxy.PercentageProxy{
				ID:        int(proxyTemp["ID"].(float64)),
				Host:      proxyTemp["Host"].(string),
				Port:      int(proxyTemp["Port"].(float64)),
				TypeName:  proxyTemp["TypeName"].(string),
				Requestor: &requestor.Requestor{},
			}

			responseParams[0] = namingService.Register(name, p2)

		case "Lookup":
			params := miopPacketRequest.Body.ReqBody.Body
			name := params[0].(string)
			responseParams[0] = namingService.Lookup(name)
		case "List":
			responseParams[0] = namingService.List()
		}

		repHeader := miop.ReplyHeader{Context: "context", RequestId: miopPacketRequest.Body.ReqHeader.RequestId, Status: 1}
		repBody := miop.ReplyBody{OperationResult: responseParams}

		header := miop.Header{Magic: "MIOP", Version: "1.0", ByteOrder: true, MessageType: 2, Size: 0}
		body := miop.Body{RepHeader: repHeader, RepBody: repBody}

		miopPacketReply = miop.Packet{Header: header, Body: body}
		marshalledReply := marshaller.Marshall(miopPacketReply)

		srh.SendMessage(marshalledReply)
		srh.Close()
	}
}
