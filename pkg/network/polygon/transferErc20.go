package polygon

import (
	"math/big"

	"github.com/MyriadFlow/superiad/generated/generc20"
	"github.com/MyriadFlow/superiad/pkg/wallet"
	rawtrasaction "github.com/MyriadFlow/superiad/pkg/wallet/rawtransaction"
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
