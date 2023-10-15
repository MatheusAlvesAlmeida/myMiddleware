package srh

type ServerRequestHandler interface {
	Server(ServerHost, ServerPort string) error
}

type TCPHandler struct {
	SRH ServerRequestHandler
}

func (th *TCPHandler) Server(ServerHost, ServerPort string) error {
	return th.SRH.Server(ServerHost, ServerPort)
}

type UDPHandler struct {
	SRH ServerRequestHandler
}

func (uh *UDPHandler) Server(ServerHost, ServerPort string) error {
	return uh.SRH.Server(ServerHost, ServerPort)
}
