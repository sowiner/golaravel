package application

import (
	"context"
	"os"

	"github.com/sowiner/golaravel/config"
	"github.com/sowiner/golaravel/pkg/helps"
	"github.com/sowiner/golaravel/server"
)

type appver string

const ver appver = appver("1.0.0")

type Application struct {
	version appver
	Config  config.IConfigure
	Srv     server.Iserver
	Log     Logger
	Ctx     context.Context
}

func (a *Application) InitProcess() {
	a.version = ver
	pwd, err := os.Getwd()
	if err != nil {
		a.Log.Panic(err)
	}
	// check project default dir struct
	for _, d := range DefaultProjectFolders {
		if !helps.CheckFolder(pwd + d) {
			err = helps.CreateFolder(pwd + d)
			if err != nil {
				a.Log.Panic(err)
			}
		}
	}
}
