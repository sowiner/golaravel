package config

type IConfigure interface{}

type configure struct {
	Debug   bool
	Address string
	Port    uint
}

func NewConfig() IConfigure {
	return new(configure)
}
