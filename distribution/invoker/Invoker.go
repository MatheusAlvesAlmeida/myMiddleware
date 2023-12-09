package invoker

import (
	"strconv"

	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/marshaller"
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/miop"
	"github.com/MatheusAlvesAlmeida/myMiddleware/infrastructure/srh"
)

type Invoker struct {
	Port int
}

func NewInvoker(port int) Invoker {
	return Invoker{Port: port}
}

func (i Invoker) Invoke() {
	srh := srh.ServerRequestHandlerTCP{ServerHost: "localhost", ServerPort: strconv.Itoa(i.Port)}
	marshaller := marshaller.Marshaller{}
	replyParams := make([]interface{}, 1)
	calculator := NewPercentageCalculatorInvoker()
	context := "Success"

	for {
		messageReceived, err := srh.ReceiveMessage()
		if err != nil {
			continue
		}

		miopPacketRequest := marshaller.Unmarshall(messageReceived)
		operation := miopPacketRequest.Body.ReqHeader.Operation

		var status int
		switch operation {
		case "GetValueOf":
			params := miopPacketRequest.Body.ReqBody.Body
			if len(params) != 2 {
				status = 101 // Logical error: Invalid number of parameters
				context = err.Error()
				break
			}
			percentage := int(params[0].(float64))
			totalValue := int(params[1].(float64))
			result, err := calculator.GetValueOf(percentage, totalValue)
			if err != nil {
				status = 101 // Logical error: Other calculation error
				context = err.Error()
				break
			}
			replyParams[0] = result
			status = 100 // Success
		case "GetPercentageOf":
			params := miopPacketRequest.Body.ReqBody.Body
			if len(params) != 2 {
				status = 101 // Logical error: Invalid number of parameters
				context = err.Error()
				break
			}
			partialValue := int(params[0].(float64))
			totalValue := int(params[1].(float64))
			result, err := calculator.GetPercentageOf(partialValue, totalValue)
			if err != nil {
				status = 101 // Logical error: Other calculation error
				context = err.Error()
				break
			}
			replyParams[0] = result
			status = 100 // Success
		default:
			status = 101 // Logical error: Unsupported operation
		}

		repHeader := miop.ReplyHeader{Context: context, RequestId: miopPacketRequest.Body.ReqHeader.RequestId, Status: status}
		repBody := miop.ReplyBody{OperationResult: replyParams}

		header := miop.Header{Magic: "MIOP", Version: "1.0", ByteOrder: true, MessageType: 2, Size: 0}
		body := miop.Body{RepHeader: repHeader, RepBody: repBody}

		miopPacketReply := miop.Packet{Header: header, Body: body}
		marshalledReply := marshaller.Marshall(miopPacketReply)

		// Sleep 10 minutes
		//time.Sleep(10 * time.Minute)

		srh.SendMessage(marshalledReply)
	}
}
