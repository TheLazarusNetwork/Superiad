package polygon

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/config/envconfig"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/wallet"
	"github.com/ethereum/go-ethereum/common"
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

func GetBalanceFromWalletAddress(walletAddress string) (*big.Int, error) {
	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		return nil, fmt.Errorf("failed to dial rpc client :%w", err)
	}
	bal, err := client.BalanceAt(context.Background(), common.HexToAddress(walletAddress), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call BalanceAt :%w", err)
	}
	return bal, nil
}

func GetBalanceInDecimalsFromWalletAddress(walletAddress string) (*big.Float, error) {
	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		return nil, fmt.Errorf("failed to dial rpc client :%w", err)
	}
	bal, err := client.BalanceAt(context.Background(), common.HexToAddress(walletAddress), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call BalanceAt :%w", err)
	}

	ether := big.NewFloat(0).SetInt(big.NewInt(1000000000000000000))
	balanceInDecimals := big.NewFloat(0).SetInt(bal)
	balanceInDecimals.Quo(balanceInDecimals, ether)
	return balanceInDecimals, nil
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

func GetWalletPrivateKey(mnemonic string) (string, error) {
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return "", err
	}
	privateKeyBytes := crypto.FromECDSA(privKey)

	return hex.EncodeToString(privateKeyBytes[2:]), nil
}
