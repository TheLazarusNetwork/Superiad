package polygon

import (
	"math/big"

	"github.com/MyriadFlow/superiad/generated/generc721"
	"github.com/MyriadFlow/superiad/pkg/wallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ERC721IsOwner(mnemonic string, contractAddr common.Address, tokenId *big.Int) (bool, error) {
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return false, err
	}
	publicKey := privKey.PublicKey
	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		return false, err
	}

	ins, err := generc721.NewErc721(contractAddr, client)
	if err != nil {
		return false, err
	}
	ownerAddr, err := ins.OwnerOf(nil, tokenId)
	if err != nil {
		return false, err
	} else if ownerAddr.String() == crypto.PubkeyToAddress(publicKey).String() {
		return true, nil
	} else {
		return false, nil
	}
}
