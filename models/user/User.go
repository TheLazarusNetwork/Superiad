package user

import (
	"log"

	"github.com/TheLazarusNetwork/mtwallet/config/dbconfig"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/wallet"
	"github.com/google/uuid"
)

type User struct {
	UserId   string `json:"userId" gorm:"primary_key"`
	Mnemonic string `json:"mnemonic" gorm:"unique;not null"`
}

var initDone bool = false

func Init() {
	if initDone {
		return
	}
	db := dbconfig.GetDb()
	if err := db.AutoMigrate(&User{}).Error; err != nil {
		log.Fatal(err)
	}
	initDone = true
}

func AddUser() (string, error) {
	Init()
	mnemonic, err := wallet.GenerateMnemonic()
	if err != nil {
		return "", err
	}
	db := dbconfig.GetDb()
	userId := uuid.NewString()
	newUser := User{
		UserId:   userId,
		Mnemonic: *mnemonic,
	}
	if err := db.Model(&newUser).Create(&newUser).Error; err != nil {
		return "", err
	} else {
		return userId, nil
	}
}

func GetMnemonic(userId string) (mnemonic string, err error) {
	Init()
	db := dbconfig.GetDb()
	var user User
	err = db.Model(&User{}).Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		return "", err
	} else {
		return user.Mnemonic, nil
	}
}
