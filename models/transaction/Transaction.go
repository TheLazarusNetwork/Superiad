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
	if initDone {
		return
	}
	db := dbconfig.GetDb()
	if err := db.AutoMigrate(&Transaction{}).Error; err != nil {
		log.Fatal(err)
	}
	initDone = true
}
