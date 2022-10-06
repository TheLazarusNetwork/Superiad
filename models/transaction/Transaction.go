package transaction

import (
	"log"

	"github.com/TheLazarusNetwork/mtwallet/config/dbconfig"
)

type Transaction struct {
	UserId         string `gorm:"not null"`
	TrasactionHash string `gorm:"primary_key"`
}

var initDone bool = false

func Init() {
	//TODO: better transaction handling
	if initDone {
		return
	}
	db := dbconfig.GetDb()
	if err := db.AutoMigrate(&Transaction{}); err != nil {
		log.Fatal(err)
	}
	initDone = true
}
