package crh

import (
	"fmt"
	"net"
)

type ClientRequestHandler struct {
	ServerHost string
	ServerPort string
	conn       net.Conn
}

func (crh *ClientRequestHandler) SendReceive(message []byte, protocol string) ([]byte, error) {
	if crh.conn == nil {
		addr := fmt.Sprintf("%s:%s", crh.ServerHost, crh.ServerPort)
		conn, err := net.Dial(protocol, addr)
		if err != nil {
			return nil, err
		}
		crh.conn = conn
	}

	_, err := crh.conn.Write(message)
	if err != nil {
		return nil, err
	}

	response := make([]byte, 1024)
	_, err = crh.conn.Read(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
