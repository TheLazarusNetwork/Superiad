package polygon

import (
	"math/big"

	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/generated/generc20"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/wallet"
	rawtrasaction "github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/wallet/rawtransaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TransferERC20(mnemonic string, toAddress common.Address, contractAddr common.Address, amount big.Int) (string, error) {
	// TODO: Ammount, check in goethbook guide
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return "", err
	}

	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		return "", err
	}
	ins, err := generc20.NewErc20(contractAddr, client)
	if err != nil {
		return "", err
	}
	decimals, err := ins.Decimals(nil)
	if err != nil {
		return "", err
	}
	decimalsCal := big.NewInt(0)
	decimalsCal.Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	amount.Mul(&amount, decimalsCal)
	chainId, err := GetChainId()
	if err != nil {
		return "", err
	}
	tx, err := rawtrasaction.SendRawTrasac(privKey, *client, int64(chainId), 310000, contractAddr, generc20.Erc20MetaData.ABI, "transfer", toAddress, &amount)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func TransferERC20AcceptFloat(mnemonic string, toAddress common.Address, contractAddr common.Address, amount float64) (string, error) {
	// TODO: Ammount, check in goethbook guide
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return "", err
	}

	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		return "", err
	}
	ins, err := generc20.NewErc20(contractAddr, client)
	if err != nil {
		return "", err
	}
	decimals, err := ins.Decimals(nil)
	if err != nil {
		return "", err
	}
	decimalsCal := big.NewInt(0)
	decimalsCal.Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	// amount.Mul(&amount, decimalsCal)
	var amountInInt64 = int64(amount * float64(10000))
	var amountInBigInt = decimalsCal.Mul(decimalsCal, big.NewInt(amountInInt64))
	amountInBigInt.Div(amountInBigInt, big.NewInt(10000))

	chainId, err := GetChainId()
	if err != nil {
		return "", err
	}
	tx, err := rawtrasaction.SendRawTrasac(privKey, *client, int64(chainId), 310000, contractAddr, generc20.Erc20MetaData.ABI, "transfer", toAddress, &amountInBigInt)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}
