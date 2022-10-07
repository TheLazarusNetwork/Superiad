package wallet

import (
	"crypto/ecdsa"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/tyler-smith/go-bip39"
)

func GenerateMnemonic() (*string, error) {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return nil, err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}
	return &mnemonic, nil
}

func GetWallet(mnemonic string, path string) (*ecdsa.PrivateKey, error) {
	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "")
	// Generate a new master node using the seed.
	var masterKey *hdkeychain.ExtendedKey
	var err error
	paths := strings.Split(path, "/")
	for _, v := range paths {
		if v == "m" {
			masterKey, err = hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
			if err != nil {
				return nil, err
			}
			continue
		}
		if strings.Contains(v, "H") {
			hSplit, err := strconv.Atoi(strings.TrimSuffix(v, "H"))
			if err != nil {
				return nil, err
			}
			masterKey, err = masterKey.Derive(uint32(hdkeychain.HardenedKeyStart + hSplit))
			if err != nil {
				return nil, err
			}
		} else {
			res, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			masterKey, err = masterKey.Derive(uint32((res)))
			if err != nil {
				return nil, err
			}
		}
	}

	btcecPrivKey, err := masterKey.ECPrivKey()
	if err != nil {
		return nil, err
	}
	privateKey := btcecPrivKey.ToECDSA()
	return privateKey, nil
}
