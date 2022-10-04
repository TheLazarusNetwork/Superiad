package user

import (
	"testing"

	"github.com/TheLazarusNetwork/mtwallet/config/envconfig"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/logwrapper"
)

func Test_AddUser(t *testing.T) {
	envconfig.InitEnvVars()
	logwrapper.Init()
	_, err := AddUser()
	if err != nil {
		t.Fatal(err)
	}
}
