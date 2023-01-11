package bsc

import (
	"context"
	"fmt"
	"math/big"

	"github.com/TheLazarusNetwork/superiad/config/envconfig"
	"github.com/TheLazarusNetwork/superiad/pkg/wallet"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetRpcUrl() string {
	return envconfig.EnvVars.NETWORK_RPC_URL_BSC
}

func GetChainId() (int, error) {
	return 97, nil // BSC testnet
}

func GetNetworkInfo() (*networkInfo, error) {
	chainId, err := GetChainId()
	if err != nil {
		return nil, err
	}
	return &networkInfo{
		Name:    "binance smart chain",
		ChainId: big.NewInt(int64(chainId)),
	}, nil
}

func GetPath() string {
	return "m/44H/60H/0H/0/0"
}

func GetBalance(mnemonic string) (*big.Int, error) {
	privkey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return nil, err
	}
	publicKey := privkey.PublicKey
	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		return nil, fmt.Errorf("failed to dial RPC client: %w", err)
	}
	bal, err := client.BalanceAt(context.Background(), crypto.PubkeyToAddress(publicKey), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %v", err)
	}
	return bal, nil
}

func GetWalletAddres(mnemonic string) (string, error) {
	privkey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return "", err
	}
	walletAddr := crypto.PubkeyToAddress(privkey.PublicKey)

	return walletAddr.String(), nil
}
