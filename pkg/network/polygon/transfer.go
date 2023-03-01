package polygon

import (
	"context"
	"math/big"

	"github.com/MyriadFlow/superiad/pkg/wallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	log "github.com/sirupsen/logrus"
)

func Transfer(mnemonic string, to common.Address, value big.Int) (string, error) {
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return "", err
	}
	publicKey := privKey.PublicKey

	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		return "", err
	}
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(publicKey))
	if err != nil {
		return "", err
	}

	maxPriorityFeePerGas, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to suggestGasTipCap")
		return "", err
	}
	chainId, err := GetChainId()
	if err != nil {
		return "", err
	}
	config := &params.ChainConfig{
		ChainID: big.NewInt(int64(chainId)),
	}
	bn, _ := client.BlockNumber(context.Background())

	bignumBn := big.NewInt(0).SetUint64(bn)
	blk, _ := client.BlockByNumber(context.Background(), bignumBn)
	baseFee := misc.CalcBaseFee(config, blk.Header())
	big2 := big.NewInt(2)
	mulRes := big.NewInt(0).Mul(baseFee, big2)
	maxFeePerGas := big.NewInt(0).Add(mulRes, maxPriorityFeePerGas)
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(int64(chainId)),
		Nonce:     nonce,
		GasFeeCap: maxFeePerGas,
		GasTipCap: maxPriorityFeePerGas,
		Gas:       21000,
		To:        &to,
		Value:     &value,
	})
	types.SignTx(tx, types.NewLondonSigner(big.NewInt(int64(chainId))), privKey)
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(int64(chainId))), privKey)
	if err != nil {
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}
