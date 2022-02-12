package models

import (
	"github.com/TheLazarusNetwork/mtwallet/models/transaction"
	"github.com/TheLazarusNetwork/mtwallet/models/user"
)

func Migrate() {
	user.Init()
	transaction.Init()
}
