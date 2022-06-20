package golaravel

import (
	"context"
	"os"

	"github.com/sowiner/golaravel/application"
	"github.com/sowiner/golaravel/config"
	"github.com/sowiner/golaravel/pkg/helps"
	"github.com/sowiner/golaravel/server"
)

type appver string

const ver appver = appver("1.0.0")

type Application struct {
	version appver
	Config  *config.Configure
	Srv     server.Iserver
	Ctx     context.Context
}

func New(opt *config.Configure) *Application {
	var cfg *config.Configure
	if opt == nil {
		cfg = opt
	} else {
		cfg = config.NewConfig()
	}
	return &Application{
		version: ver,
		Config:  cfg,
		Ctx:     context.Background(),
		Srv:     server.NewServer,
	}
}

func (a *Application) InitProcess() {
	pwd, err := os.Getwd()
	if err != nil {
		a.Config.Log.Panic(err)
	}

	a.Config.Log.Info(pwd)
	// check project default dir struct
	for _, d := range application.DefaultProjectFolders {
		if !helps.CheckFolder(pwd + d) {
			err = helps.CreateFolder(pwd + d)
			if err != nil {
				a.Config.Log.Panic(err)
			}
		}
	}
}
