package user

import (
	"log"

	"github.com/TheLazarusNetwork/mtwallet/config/dbconfig"
	"github.com/TheLazarusNetwork/mtwallet/models/transaction"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/wallet"
	"github.com/google/uuid"
)

type User struct {
	UserId       string                    `json:"userId" gorm:"primary_key"`
	Mnemonic     string                    `json:"-" gorm:"unique;not null"`
	Transactions []transaction.Transaction `json:"transactions" gorm:"foreignKey:UserId"`
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

func AddTrasactionHash(userId string, hash string) error {
	db := dbconfig.GetDb()
	trx := transaction.Transaction{UserId: userId, TrasactionHash: hash}
	association := db.Model(&User{UserId: userId}).Association("Transactions")
	err := association.Error
	if err != nil {
		return err
	}
	err = association.Append(&trx).Error
	if err != nil {
		return err
	}
	return nil
}

func GetMnemonic(userId string) (mnemonic string, err error) {
	user, err := GetUser(userId)
	if err != nil {
		return "", err
	} else {
		return user.Mnemonic, nil
	}
}

func GetUser(userId string) (user User, err error) {
	Init()
	db := dbconfig.GetDb()
	err = db.Model(&User{}).Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		return User{}, err
	} else {
		return user, nil
	}
}
