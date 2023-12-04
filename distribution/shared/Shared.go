package shared

const SERVER_PORT = "8080"
const NAMING_PORT = 8081

type Invocation struct {
	Host    string
	Port    int
	Request Request
}

type Termination struct {
	Rep Reply
}

type IOR struct {
	Host string
	Port int
	Id   int
}

type Request struct {
	Op     string
	Params []interface{}
}

type Reply struct {
	Result []interface{}
}
