package crh

import (
	"net"
	"time"

	errorhandler "github.com/MatheusAlvesAlmeida/myMiddleware/distribution/error_handler" // Import the error handler package
)

type ClientRequestHandlerTCP struct {
	ServerAddress string
	Conn          net.Conn
	ErrorHandler  errorhandler.ErrorHandler // Include the error handler
}

func NewClientRequestHandlerTCP(serverAddress string) *ClientRequestHandlerTCP {
	return &ClientRequestHandlerTCP{
		ServerAddress: serverAddress,
		ErrorHandler:  errorhandler.ErrorHandler{}, // Instantiate the error handler
	}
}

func (crh *ClientRequestHandlerTCP) establishConnection() error {
	conn, err := net.Dial("tcp", crh.ServerAddress)
	if err != nil {
		return err
	}
	crh.Conn = conn
	return nil
}

func (crh *ClientRequestHandlerTCP) SendReceive(message []byte) ([]byte, error) {
	if crh.Conn == nil {
		err := crh.establishConnection()
		if err != nil {
			return nil, err
		}
	}

	_, err := crh.Conn.Write(message)
	if err != nil {
		err = crh.ErrorHandler.HandleError(&crh.Conn, err)
		if err != nil {
			return nil, err
		}
		return crh.SendReceive(message) // Retry handled by error handler
	}

	response := make([]byte, 1024)
	crh.Conn.SetReadDeadline(time.Now().Add(5 * time.Second)) // Set a read deadline
	n, err := crh.Conn.Read(response)
	if err != nil {
		err = crh.ErrorHandler.HandleError(&crh.Conn, err)
		if err != nil {
			return nil, err
		}
		return crh.SendReceive(message) // Retry handled by error handler
	}

	return response[:n], nil
}
