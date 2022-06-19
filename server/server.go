package server

type Iserver interface{}

type server struct{}

func NewServer() Iserver {
	return new(server)
}
