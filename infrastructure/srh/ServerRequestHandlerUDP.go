package srh

import (
	"fmt"
	"net"
)

type ServerRequestHandlerUDP struct {
	ServerHost string
	ServerPort string
	clientAddr *net.UDPAddr
	Conn       *net.UDPConn
}

func (srh *ServerRequestHandlerUDP) Server(ServerHost, ServerPort string) *ServerRequestHandlerUDP {
	response := new(ServerRequestHandlerUDP)
	response.ServerHost = ServerHost
	response.ServerPort = ServerPort
	response.clientAddr = nil
	response.Conn = nil

	return response
}

func (srh *ServerRequestHandlerUDP) ReceiveMessage() ([]byte, *net.UDPAddr, error) {
	if srh.Conn == nil {
		udpAddr, err := net.ResolveUDPAddr("udp", srh.ServerHost+":"+srh.ServerPort)
		if err != nil {
			return nil, nil, err
		}

		conn, err := net.ListenUDP("udp", udpAddr)
		if err != nil {
			return nil, nil, err
		}

		srh.Conn = conn
	}

	buffer := make([]byte, 1024)
	_, remoteAddr, err := srh.Conn.ReadFromUDP(buffer)
	if err != nil {
		return nil, nil, err
	}

	return buffer, remoteAddr, nil
}

func (srh *ServerRequestHandlerUDP) SendMessage(clientAddr *net.UDPAddr, message []byte) error {
	if srh.Conn == nil {
		return fmt.Errorf("connection is nil")
	}

	_, err := srh.Conn.WriteToUDP(message, clientAddr)
	if err != nil {
		fmt.Printf("Error sending response: %s\n", err)
		return err
	}

	return nil
}
