package polygon

import (
	"github.com/TheLazarusNetwork/mtwallet/generated/generc721"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/wallet"
	rawtrasaction "github.com/TheLazarusNetwork/mtwallet/util/pkg/wallet/rawtransaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SetAprovalForAllErc721(mnemonic string, operatorAddr common.Address, contractAddr common.Address, approved bool) (string, error) {
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return "", err
	}

	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		return "", err
	}
	chainId, err := GetChainId()
	if err != nil {
		return "", err
	}
	tx, err := rawtrasaction.SendRawTrasac(privKey, *client, int64(chainId), 310000, contractAddr, generc721.Erc721MetaData.ABI, "setApprovalForAll", operatorAddr, &approved)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil

}
