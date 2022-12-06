package wallet

import (
	"context"
	"errors"

	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/network/polygon"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/store"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/wallet"
	"github.com/google/uuid"
)

var (
	ErrNoRecordFound = errors.New("no record found")
)

type Wallet struct {
	Address    string `json:"address"`
	Salt       string `json:"salt"`
	Mnemonic   string `json:"mnemonic"`
	PrivateKey string `json:"privateKey"`
}

func GenerateWallet() (Wallet, error) {
	// Generate Mnemonic and derive private/public keys
	mnemonic, err := wallet.GenerateMnemonic()
	if err != nil {
		logo.Fatal("failed to generate wallet", err)
		return Wallet{}, err
	}
	privKey, err := polygon.GetWalletPrivateKey(*mnemonic)
	if err != nil {
		logo.Fatal("failed to derive wallet private key", err)
		return Wallet{}, err
	}
	walletAddress, err := polygon.GetWalletAddres(*mnemonic)
	if err != nil {
		logo.Fatal("failed to generate wallet", err)
		return Wallet{}, err
	}
	salt := uuid.NewString()

	err = store.RDS.Set(context.Background(), walletAddress, salt, 0).Err()
	if err != nil {
		logo.Fatal("failed to generate wallet", err)
		return Wallet{}, err
	}

	newWallet := Wallet{
		Address:    walletAddress,
		Salt:       salt,
		Mnemonic:   *mnemonic,
		PrivateKey: privKey,
	}
	return newWallet, nil
}

func GetWallet(walletAddress string) (string, error) {
	// Find the wallet address based on the salt
	salt, err := store.RDS.Get(context.Background(), walletAddress).Result()
	if err != nil {
		logo.Errorf("failed to query wallet address from redis, error: %s", err)
		return "", err
	}
	return salt, nil
}
