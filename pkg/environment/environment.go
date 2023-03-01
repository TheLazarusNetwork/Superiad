// Package environment provides way to define and get environment
package environment

import (
	"log"

	"github.com/MyriadFlow/superiad/config/envconfig"
)

type Environment int

const (
	PROD Environment = iota
	DEV  Environment = iota
)

// GetEnvironment returns environment from env variable and fatals if it is different from
// PROD and DEV
func GetEnvironment() Environment {
	appEnv := envconfig.EnvVars.APP_ENVIRONMENT

	if appEnv == "PROD" {
		return PROD
	} else if appEnv == "DEV" {
		return DEV
	} else {
		log.Fatal("App environment not supported")
		return -1
	}
}
