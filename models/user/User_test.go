package user

import (
	"testing"

	"github.com/TheLazarusNetwork/mtwallet/config"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/logwrapper"
)

func Test_AddUser(t *testing.T) {
	config.Init("../../.env")
	logwrapper.Init("../..", false)
	_, err := AddUser()
	if err != nil {
		t.Fatal(err)
	}
}
