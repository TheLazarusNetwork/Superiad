// Package dbconinit provides method to Init database
package dbconinit

import (
	"fmt"

	"github.com/MyriadFlow/superiad/config/envconfig"
	"github.com/MyriadFlow/superiad/pkg/store"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() {
	var (
		host     = envconfig.EnvVars.DB_HOST
		username = envconfig.EnvVars.DB_USERNAME
		password = envconfig.EnvVars.DB_PASSWORD
		dbname   = envconfig.EnvVars.DB_NAME
		port     = envconfig.EnvVars.DB_PORT
	)

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%d",
		host, username, password, dbname, port)

	var err error
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dns,
	}))
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("failed to Create database config", err)
	}

	// Store database in global store
	store.DB = db

	// Get underlying sql database to ping it
	sqlDb, err := db.DB()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("failed to Connect to database", err)
	}

	// If ping fails then log error and exit
	if err = sqlDb.Ping(); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("failed to ping database", err)
	}
}
