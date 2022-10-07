// Package dbmigrate provides method to migrate models into database
package dbmigrate

import (
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/TheLazarusNetwork/mtwallet/models/transaction"
	"github.com/TheLazarusNetwork/mtwallet/models/user"
	"github.com/TheLazarusNetwork/mtwallet/pkg/store"
)

func Migrate() {
	db := store.DB
	err := db.AutoMigrate(&user.User{}, &transaction.Transaction{})
	if err != nil {
		logo.Fatalf("failed to migrate models into database: %s", err)
	}
}
