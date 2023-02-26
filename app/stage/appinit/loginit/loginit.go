// Package loginit provides method to Init loging config
package loginit

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var StandardFields = log.Fields{
	"hostname": "host-server",
	"appname":  "sotreus",
}

func Init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)

	// Get Hostname for updating Log StandardFields
	HostName, err := os.Hostname()
	if err != nil {
		log.Infof("Error in getting the Hostname: %v", err)
	} else {
		StandardFields = log.Fields{
			"hostname": HostName,
			"appname":  "Erebrus",
		}
	}
}
