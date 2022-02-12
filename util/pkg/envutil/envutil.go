package envutil

import (
	"os"

	"github.com/TheLazarusNetwork/mtwallet/util/pkg/logwrapper"
)

func MustGetEnv(key string) string {
	logwrapper.Init("", false)
	val := os.Getenv(key)
	if val == "" {
		logwrapper.Fatalf("env variable %v is not defined", key)
	}
	return val
}
