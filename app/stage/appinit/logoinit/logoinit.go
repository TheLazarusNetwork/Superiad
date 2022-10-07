// Package logoinit provides method to Init loging config
package logoinit

import (
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/TheLazarusNetwork/mtwallet/pkg/environment"
	"github.com/sirupsen/logrus"
)

func Init() {
	logrusEntry := logrus.New().WithFields(logrus.Fields{})

	if environment.GetEnvironment() == environment.PROD {
		logrusEntry.Logger.SetFormatter(&logrus.JSONFormatter{})
	}

	logo.SetInstance(*logrusEntry)
}
