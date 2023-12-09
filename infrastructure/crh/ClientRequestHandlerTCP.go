package crh

import (
	"fmt"
	"net"
	"time"

	errorhandler "github.com/MatheusAlvesAlmeida/myMiddleware/distribution/error_handler"
)

type ClientRequestHandlerTCP struct {
	ServerAddress string
	Conn          net.Conn
	ErrorHandler  errorhandler.ErrorHandler
}

func NewClientRequestHandlerTCP(serverAddress string) *ClientRequestHandlerTCP {
	return &ClientRequestHandlerTCP{
		ServerAddress: serverAddress,
		ErrorHandler:  errorhandler.ErrorHandler{},
	}
}

func (crh *ClientRequestHandlerTCP) establishConnection() error {
	for {
		conn, err := net.Dial("tcp", crh.ServerAddress)
		if err != nil {
			fmt.Println("Error connecting to server: ", err)
			err = crh.ErrorHandler.HandleConnectionError(&crh.Conn, err)
			if err != nil && err.Error() == "connection timeout: retrying" {
				continue // Retry connection establishment
			}
			return err // Return the error if retries are exhausted or other errors occur
		}
		crh.Conn = conn
		return nil // Successful connection establishment
	}
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
	crh.Conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	n, err := crh.Conn.Read(response)
	if err != nil {
		err = crh.ErrorHandler.HandleError(&crh.Conn, err)
		if err != nil {
			fmt.Println("Error reading from connection: ", err)
		}

		return crh.SendReceive(message) // Retry handled by error handler
	}

	return response[:n], nil
}
