package dbconfig

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/TheLazarusNetwork/mtwallet/util/pkg/envutil"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// Return singleton instance of db, initiates it before if it is not initiated already
func GetDb() *gorm.DB {
	if db != nil {
		return db
	}
	var (
		host     = envutil.MustGetEnv("DB_HOST")
		username = envutil.MustGetEnv("DB_USERNAME")
		password = envutil.MustGetEnv("DB_PASSWORD")
		dbname   = envutil.MustGetEnv("DB_NAME")
		port     = envutil.MustGetEnv("DB_PORT")
	)

	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s",
		host, username, password, dbname, port)
	var err error
	db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	if err = db.DB().Ping(); err != nil {
		log.Fatal("failed to ping database", err)
	}
	return db
}
