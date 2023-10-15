package crh

import (
	"net"
)

type ClientRequestHandlerUDP struct {
	ServerAddress string
	Conn          *net.UDPConn
}

func (crh *ClientRequestHandlerUDP) SendReceive(message []byte) ([]byte, error) {
	if crh.Conn == nil {
		udpAddr, err := net.ResolveUDPAddr("udp", crh.ServerAddress)
		if err != nil {
			return nil, err
		}

		conn, err := net.DialUDP("udp", nil, udpAddr)
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
	n, _, err := crh.Conn.ReadFromUDP(response)
	if err != nil {
		return nil, err
	}

	return response[:n], nil
}
