package srh

import (
	"fmt"
	"net"
)

type ServerRequestHandlerTCP struct {
	Protocol string
}

func (srh *ServerRequestHandlerTCP) Server(ServerHost, ServerPort string) error {
	listenAddr := ServerHost + ":" + ServerPort
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	defer listener.Close()

	fmt.Printf("Server listening on %s (TCP)...\n", listenAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %s\n", err)
			continue
		}

		go srh.handleConnection(conn)
	}
}

func (srh *ServerRequestHandlerTCP) handleConnection(conn net.Conn) {
	defer conn.Close()

	receivedMessage := srh.receiveMessage(conn)
	srh.sendMessage(conn, receivedMessage)
}

func (srh *ServerRequestHandlerTCP) receiveMessage(conn net.Conn) []byte {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("Error reading data: %s\n", err)
		return nil
	}
	return buffer[:n]
}

func (srh *ServerRequestHandlerTCP) sendMessage(conn net.Conn, message []byte) {
	fmt.Printf("Debug info - Received message: %s\n", string(message))
	fromServerText := "From server: "
	message = append([]byte(fromServerText), message...)
	_, err := conn.Write(message)
	if err != nil {
		fmt.Printf("Error sending response: %s\n", err)
	}
}
