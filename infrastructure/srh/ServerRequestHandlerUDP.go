package srh

import (
	"fmt"
	"net"
)

type ServerRequestHandlerUDP struct {
	Protocol string
}

func (srh *ServerRequestHandlerUDP) Server(ServerHost, ServerPort string) error {
	listenAddr := ServerHost + ":" + ServerPort
	udpAddr, err := net.ResolveUDPAddr("udp", listenAddr)
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Printf("Server listening on %s (UDP)...\n", listenAddr)

	for {
		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error reading data: %s\n", err)
			continue
		}

		go srh.handleConnection(conn, addr, buffer[:n])
	}
}

func (srh *ServerRequestHandlerUDP) handleConnection(conn *net.UDPConn, addr *net.UDPAddr, data []byte) {
	receivedMessage := srh.receiveMessage(data)
	srh.sendMessage(conn, addr, receivedMessage)
}

func (srh *ServerRequestHandlerUDP) receiveMessage(data []byte) []byte {
	return data
}

func (srh *ServerRequestHandlerUDP) sendMessage(conn *net.UDPConn, addr *net.UDPAddr, message []byte) {
	fmt.Printf("Debug info - Received message: %s\n", string(message))
	_, err := conn.WriteTo(message, addr)
	if err != nil {
		fmt.Printf("Error sending response: %s\n", err)
	}
}
