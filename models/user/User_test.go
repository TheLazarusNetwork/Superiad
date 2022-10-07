package user

import (
	"testing"

	"github.com/TheLazarusNetwork/mtwallet/app/stage/appinit"
)

func Test_AddUser(t *testing.T) {
	appinit.Init()
	_, err := AddUser()
	if err != nil {
		t.Fatal(err)
	}
}
