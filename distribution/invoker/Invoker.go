package invoker

import (
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/marshaller"
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/miop"
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/shared"
	"github.com/MatheusAlvesAlmeida/myMiddleware/infrastructure/srh"
)

type Invoker struct {
	Port int
}

func NewInvoker() Invoker {
	return Invoker{}
}

func (Invoker) Invoke() {
	srh := srh.ServerRequestHandlerTCP{ServerHost: "localhost", ServerPort: shared.SERVER_PORT}
	marshaller := marshaller.Marshaller{}
	replyParams := make([]interface{}, 1)
	calculator := NewPercentageCalculatorInvoker()

	for {
		messageReceived, err := srh.ReceiveMessage()
		if err != nil {
			continue
		}

		miopPacketRequest := marshaller.Unmarshall(messageReceived)
		operation := miopPacketRequest.Body.ReqHeader.Operation

		switch operation {
		case "GetValueOf":
			params := miopPacketRequest.Body.ReqBody.Body
			percentage := int(params[0].(float64))
			totalValue := int(params[1].(float64))
			replyParams[0] = calculator.GetValueOf(percentage, totalValue)
		case "GetPercentageOf":
			params := miopPacketRequest.Body.ReqBody.Body
			partialValue := int(params[0].(float64))
			totalValue := int(params[1].(float64))
			replyParams[0] = calculator.GetPercentageOf(partialValue, totalValue)
		}

		repHeader := miop.ReplyHeader{Context: "context", RequestId: miopPacketRequest.Body.ReqHeader.RequestId, Status: 1}
		repBody := miop.ReplyBody{OperationResult: replyParams}

		header := miop.Header{Magic: "MIOP", Version: "1.0", ByteOrder: true, MessageType: 2, Size: 0}
		body := miop.Body{RepHeader: repHeader, RepBody: repBody}

		miopPacketReply := miop.Packet{Header: header, Body: body}
		marshalledReply := marshaller.Marshall(miopPacketReply)

		srh.SendMessage(marshalledReply)
	}

}
