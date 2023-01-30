package polygon

import (
	"fmt"

	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/config/envconfig"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/generated/virtuacoinnft"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/wallet"
	rawtrasaction "github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/wallet/rawtransaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DelegateErc721(walletAddress string, contractAddr common.Address, metadataURI string) (string, error) {
	operatorPrivKey, err := wallet.GetWallet(envconfig.EnvVars.OPERATOR_MNEMONIC, GetPath())
	if err != nil {
		err = fmt.Errorf("failed to get operator wallet from mnemonic: %w", err)
		return "", err
	}

	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		err = fmt.Errorf("failed to dial rpc url: %w", err)
		return "", err
	}
	chainId, err := GetChainId()
	if err != nil {
		return "", err
	}

	creatorAddress := common.HexToAddress(walletAddress)

	tx, err := rawtrasaction.SendRawTrasac(operatorPrivKey, *client, int64(chainId), 310000, contractAddr, virtuacoinnft.VirtuacoinnftABI, "delegateAssetCreation", creatorAddress, metadataURI)
	if err != nil {
		err = fmt.Errorf("failed to send raw transaction: %w", err)
		return "", err
	}
	return tx.Hash().Hex(), nil

}
