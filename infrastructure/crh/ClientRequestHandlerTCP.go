package crh

import (
	"net"
)

type ClientRequestHandlerTCP struct {
	ServerAddress string
	Conn          net.Conn
}

func (crh *ClientRequestHandlerTCP) SendReceive(message []byte) ([]byte, error) {
	if crh.Conn == nil {
		conn, err := net.Dial("tcp", crh.ServerAddress)
		if err != nil {
			return nil, err
		}
		crh.Conn = conn
	}

	_, err := crh.Conn.Write(message)
	if err != nil {
		return nil, err
	}

	response := make([]byte, 1024)
	n, err := crh.Conn.Read(response)
	if err != nil {
		return nil, err
	}

	return response[:n], nil
}
