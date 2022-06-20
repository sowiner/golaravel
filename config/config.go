package config

import (
	"github.com/sirupsen/logrus"
	"github.com/sowiner/golaravel/logger"
)

type Configure struct {
	Log     logger.Logger
	Debug   bool
	Address string
	Port    uint
}

func NewConfig() *Configure {
	c := &Configure{
		Log:     logrus.New(),
		Debug:   false,
		Address: "127.0.0.1",
		Port:    8000,
	}
	return c
}
