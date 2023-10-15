package srh

import (
	"fmt"
	"net"
)

type ServerRequestHandlerTCP struct {
	ServerHost string
	ServerPort string
	Conn       net.Conn
}

func (srh *ServerRequestHandlerTCP) Server(ServerHost, ServerPort string) *ServerRequestHandlerTCP {
	response := new(ServerRequestHandlerTCP)
	response.ServerHost = ServerHost
	response.ServerPort = ServerPort
	response.Conn = nil

	return response
}

func (srh *ServerRequestHandlerTCP) ReceiveMessage() ([]byte, error) {
	listenAddr := srh.ServerHost + ":" + srh.ServerPort
	listener, err := net.Listen("tcp", listenAddr)

	if err != nil {
		return nil, err
	}
	defer listener.Close()

	conn, err := listener.Accept()

	if err != nil {
		return nil, err
	} else {
		srh.Conn = conn
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)

	if err != nil {
		return nil, err
	}

	return buffer[:n], nil
}

func (srh *ServerRequestHandlerTCP) SendMessage(message []byte) {
	if srh.Conn == nil {
		return
	}

	_, err := srh.Conn.Write(message)
	if err != nil {
		fmt.Printf("Error sending response: %s\n", err)
	}
}
