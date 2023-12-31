package requestor

import (
	"strconv"

	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/interceptor"
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/marshaller"
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/miop"
	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/shared"
	"github.com/MatheusAlvesAlmeida/myMiddleware/infrastructure/crh"
)

type Requestor struct {
	ClientRequestHandler *crh.ClientRequestHandlerTCP
}

func NewRequestor() Requestor {
	return Requestor{ClientRequestHandler: nil}
}

func __mountRequestPacket(invoker shared.Invocation) miop.Packet {
	reqHeader := miop.RequestHeader{Context: invoker.Context, RequestId: 1000, ResponseExpected: true, ObjectKey: 2000, Operation: invoker.Request.Op}
	reqBody := miop.RequestBody{Body: invoker.Request.Params}
	header := miop.Header{Magic: "MIOP", Version: "1.0", ByteOrder: true, MessageType: 1, Size: 1024}
	body := miop.Body{ReqHeader: reqHeader, ReqBody: reqBody}
	miopPacketRequest := miop.Packet{Header: header, Body: body}

	return miopPacketRequest
}

func (r *Requestor) Invoke(invoker shared.Invocation) interface{} {
	interceptor := interceptor.NewInvocationInterceptor()

	if r.ClientRequestHandler == nil {
		serverAddress := invoker.Host + ":" + strconv.Itoa(invoker.Port)
		r.ClientRequestHandler = &crh.ClientRequestHandlerTCP{ServerAddress: serverAddress}
	}

	marshaller := marshaller.Marshaller{}
	miopPacketRequest := __mountRequestPacket(invoker)
	interceptor.Intercept(miopPacketRequest, true, false)

	msgToClientBytes := marshaller.Marshall(miopPacketRequest)
	interceptor.Intercept(miopPacketRequest, false, true)

	msgFromServerBytes, err := r.ClientRequestHandler.SendReceive(msgToClientBytes)
	if err != nil {
		panic(err)
	}
	miopPacketReply := marshaller.Unmarshall(msgFromServerBytes)
	interceptor.Intercept(miopPacketReply, false, false)

	if miopPacketReply.Body.RepHeader.Status != 100 {
		errMessage := miopPacketReply.Body.RepHeader.Context
		panic(errMessage)
	}

	response := miopPacketReply.Body.RepBody.OperationResult

	return response
}
