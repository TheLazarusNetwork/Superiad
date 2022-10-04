package main

import (
	"fmt"

	"github.com/TheLazarusNetwork/mtwallet/app"
	"github.com/TheLazarusNetwork/mtwallet/config/envconfig"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/logwrapper"
)

func main() {

	app.Init()
	logwrapper.Log.Info("Starting app")
	port := envconfig.EnvVars.APP_PORT
	err := app.GinApp.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		logwrapper.Log.Fatal(err)
	}
}
