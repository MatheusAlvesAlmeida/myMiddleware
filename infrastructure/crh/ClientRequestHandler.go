package crh

import (
	"fmt"
	"net"
)

type ClientRequestHandler struct {
	ServerHost string
	ServerPort string
	Conn       net.Conn
}

func (crh *ClientRequestHandler) SendReceive(message []byte, protocol string) ([]byte, error) {
	if crh.Conn == nil {
		addr := fmt.Sprintf("%s:%s", crh.ServerHost, crh.ServerPort)
		conn, err := net.Dial(protocol, addr)
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
	_, err = crh.Conn.Read(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
