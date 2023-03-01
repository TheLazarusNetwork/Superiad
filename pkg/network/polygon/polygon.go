package polygon

import (
	"context"
	"fmt"
	"math/big"

	"github.com/MyriadFlow/superiad/config/envconfig"
	"github.com/MyriadFlow/superiad/pkg/wallet"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetChainId() (int, error) {
	return 80001, nil
}

func GetPath() string {
	return "m/44H/60H/0H/0/0"
}

func GetRpcUrl() string {
	return envconfig.EnvVars.NETWORK_RPC_URL_POLYGON
}

func GetBalance(mnemonic string) (*big.Int, error) {
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return nil, err
	}
	publicKey := privKey.PublicKey
	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		return nil, fmt.Errorf("failed to dial rpc client :%w", err)
	}
	bal, err := client.BalanceAt(context.Background(), crypto.PubkeyToAddress(publicKey), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call BalanceAt :%w", err)
	}
	return bal, nil
}

func GetNetworkInfo() (*networkInfo, error) {
	chainId, err := GetChainId()
	if err != nil {
		return nil, err
	}
	return &networkInfo{
		Name:    "polygon",
		ChainId: big.NewInt(int64(chainId)),
	}, nil
}

func GetWalletAddres(mnemonic string) (string, error) {
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return "", err
	}
	walletAddr := crypto.PubkeyToAddress(privKey.PublicKey)

	return walletAddr.String(), nil
}
