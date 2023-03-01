// Package dbmigrate provides method to migrate models into database
package dbmigrate

import (
	"github.com/MyriadFlow/superiad/models/transaction"
	"github.com/MyriadFlow/superiad/models/user"
	"github.com/MyriadFlow/superiad/pkg/store"
	log "github.com/sirupsen/logrus"
)

func Migrate() {
	db := store.DB
	err := db.AutoMigrate(&user.User{}, &transaction.Transaction{})
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("failed to migrate models into database: %s", err)
	}
}
