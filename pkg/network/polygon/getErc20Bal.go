package polygon

import (
	"math/big"

	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/generated/generc20"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/wallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetERC20Balance(mnemonic string, contractAddr common.Address) (*big.Int, error) {
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return nil, err
	}
	publicKey := privKey.PublicKey

	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		return nil, err
	}
	ins, err := generc20.NewErc20(contractAddr, client)
	if err != nil {
		return nil, err
	}
	bal, err := ins.BalanceOf(nil, crypto.PubkeyToAddress(publicKey))
	if err != nil {
		return nil, err
	} else {
		return bal, nil
	}
}

func GetERC20BalanceFromWalletAddress(walletAddress common.Address, contractAddr common.Address) (*big.Int, error) {
	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		return nil, err
	}
	ins, err := generc20.NewErc20(contractAddr, client)
	if err != nil {
		return nil, err
	}
	bal, err := ins.BalanceOf(nil, walletAddress)
	if err != nil {
		return nil, err
	} else {
		return bal, nil
	}
}

func GetERC20BalanceInDecimalsFromWalletAddress(walletAddress common.Address, contractAddr common.Address) (*big.Float, error) {
	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		return nil, err
	}
	ins, err := generc20.NewErc20(contractAddr, client)
	if err != nil {
		return nil, err
	}
	decimals, err := ins.Decimals(nil)
	if err != nil {
		return nil, err
	}
	decimalsCal := big.NewInt(0)
	decimalsCal.Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	// amount.Mul(&amount, decimalsCal)
	bal, err := ins.BalanceOf(nil, walletAddress)
	if err != nil {
		return nil, err
	}
	token := big.NewFloat(0).SetInt(big.NewInt(decimalsCal.Int64()))
	balanceInDecimals := big.NewFloat(0).SetInt(bal)
	balanceInDecimals.Quo(balanceInDecimals, token)
	return balanceInDecimals, nil
}
