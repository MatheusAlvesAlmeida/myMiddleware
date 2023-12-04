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
	var conn net.Conn
	if srh.Conn == nil {
		listenAddr := srh.ServerHost + ":" + srh.ServerPort
		listener, err := net.Listen("tcp", listenAddr)

		if err != nil {
			return nil, err
		}
		defer listener.Close()

		fmt.Println("Waiting for connections...")

		for {
			conn, err = listener.Accept()

			fmt.Printf("Received connection from %s\n", conn.RemoteAddr().String())

			if err != nil {
				fmt.Printf("Error accepting connection: %s\n", err)
				continue
			} else {
				srh.Conn = conn
				break
			}
		}
	} else {
		conn = srh.Conn
		fmt.Println("Waiting for messages...")
	}

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Printf("Connection Lost! Log error: %s\n", err)
			srh.Conn = nil
			return nil, err
		}

		if n > 0 {
			return buffer[:n], nil
		}
	}
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
