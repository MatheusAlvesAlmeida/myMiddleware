package marshaller

import (
	"encoding/json"

	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/miop"
)

type Marshaller struct{}

func (m *Marshaller) Marshall(message miop.Packet) []byte {
	r, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	return r
}

func (m *Marshaller) Unmarshall(message []byte) miop.Packet {
	r := miop.Packet{}
	err := json.Unmarshal(message, &r)
	if err != nil {
		panic(err)
	}
	return r
}
