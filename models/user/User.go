package user

import (
	"errors"

	"github.com/MyriadFlow/superiad/models/transaction"
	"github.com/MyriadFlow/superiad/pkg/store"
	"github.com/MyriadFlow/superiad/pkg/wallet"
)

var (
	ErrNoRecordFound = errors.New("no record found")
)

type User struct {
	UserId       string                    `json:"userId" gorm:"primary_key;unique;not null"`
	Mnemonic     string                    `json:"-" gorm:"unique;not null"`
	Transactions []transaction.Transaction `json:"transactions" gorm:"foreignKey:UserId"`
	Email        string                    `json:"email" gorm:"unique;not null"`
}

func AddUser(userId string, email string) error {
	mnemonic, err := wallet.GenerateMnemonic()
	if err != nil {
		return err
	}
	db := store.DB

	newUser := User{
		UserId:   userId,
		Mnemonic: *mnemonic,
		Email:    email,
	}
	if err := db.Model(&newUser).Create(&newUser).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func AddTrasactionHash(userId string, hash string) error {
	db := store.DB
	trx := transaction.Transaction{UserId: userId, TrasactionHash: hash}
	association := db.Model(&User{UserId: userId}).Association("Transactions")
	err := association.Error
	if err != nil {
		return err
	}
	err = association.Append(&trx)
	if err != nil {
		return err
	}
	return nil
}

// func SetLockStatus(userId string, lockStatus bool) error {
// 	db := store.DB
// 	res := db.Model(&User{}).Where("user_id = ?", userId).Update("is_user_locked", lockStatus)
// 	if err := res.Error; err != nil {
// 		return err
// 	}
// 	if res.RowsAffected == 0 {
// 		return ErrNoRecordFound
// 	}

// 	return nil
// }

func GetMnemonic(userId string) (mnemonic string, err error) {
	user, err := GetUser(userId)
	if err != nil {
		return "", err
	} else {
		return user.Mnemonic, nil
	}
}

func GetUser(email string) (user User, err error) {
	db := store.DB
	err = db.Model(&User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		return User{}, err
	} else {
		return user, nil
	}
}
