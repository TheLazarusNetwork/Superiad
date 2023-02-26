package rawtrasaction

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	log "github.com/sirupsen/logrus"
)

func SendRawTrasac(privateKey *ecdsa.PrivateKey, client ethclient.Client, chainId int64, gas uint64, contractAddress common.Address, abiS string, method string, args ...interface{}) (*types.Transaction, error) {

	abiP, err := abi.JSON(strings.NewReader(abiS))
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to parse JSON abi")
		return nil, err
	}

	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Warnf("failed to get nonce")
		return nil, err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to call client.NetworkID")
		return nil, err
	}

	bytesData, err := abiP.Pack(method, args...)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to pack trasaction of method %v", method)
		return nil, err
	}

	maxPriorityFeePerGas, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to suggestGasTipCap")
		return nil, err
	}
	config := &params.ChainConfig{
		ChainID: big.NewInt(chainId),
	}
	bn, _ := client.BlockNumber(context.Background())

	bignumBn := big.NewInt(0).SetUint64(bn)
	blk, _ := client.BlockByNumber(context.Background(), bignumBn)
	baseFee := misc.CalcBaseFee(config, blk.Header())
	big2 := big.NewInt(2)
	mulRes := big.NewInt(0).Mul(baseFee, big2)
	maxFeePerGas := big.NewInt(0).Add(mulRes, maxPriorityFeePerGas)
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasFeeCap: maxFeePerGas,
		GasTipCap: maxPriorityFeePerGas,
		Gas:       gas,
		To:        &contractAddress,
		Data:      bytesData,
	})
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to sign trasaction %v, error: %s", tx, err)
		return nil, err
	}

	err = client.SendTransaction(context.TODO(), signedTx)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to send trasaction")
		return nil, err
	}
	return signedTx, nil
}
