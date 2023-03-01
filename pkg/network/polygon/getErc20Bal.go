package polygon

import (
	"math/big"

	"github.com/MyriadFlow/superiad/generated/generc20"
	"github.com/MyriadFlow/superiad/pkg/wallet"
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
