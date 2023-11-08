package marshaller

import (
	"encoding/json"
	"fmt"

	"github.com/MatheusAlvesAlmeida/myMiddleware/distribution/miop"
)

type Marshaller struct{}

func (m *Marshaller) Marshall(message miop.Packet) []byte {
	r, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Debug info: marshall error")
		panic(err)
	}
	return r
}

func (m *Marshaller) Unmarshall(message []byte) miop.Packet {
	r := miop.Packet{}
	err := json.Unmarshal(message, &r)
	if err != nil {
		fmt.Println("Debug info: unmarshall error")
		panic(err)
	}
	return r
}
