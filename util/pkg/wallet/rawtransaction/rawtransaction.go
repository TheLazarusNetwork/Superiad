package rawtrasaction

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/TheLazarusNetwork/mtwallet/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

func SendRawTrasac(privateKey *ecdsa.PrivateKey, client ethclient.Client, chainId int64, gas uint64, contractAddress common.Address, abiS string, method string, args ...interface{}) (*types.Transaction, error) {

	abiP, err := abi.JSON(strings.NewReader(abiS))
	if err != nil {
		logwrapper.Errorf("failed to parse JSON abi, error %s", err)
		return nil, err
	}

	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		logwrapper.Warnf("failed to get nonce")
		return nil, err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		logwrapper.Errorf("failed to call client.NetworkID, error: %s", err)
		return nil, err
	}

	bytesData, err := abiP.Pack(method, args...)
	if err != nil {
		logwrapper.Errorf("failed to pack trasaction of method %v, error: %s", method, err)
		return nil, err
	}

	maxPriorityFeePerGas, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		logwrapper.Errorf("failed to suggestGasTipCap, error %s", err)
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
		logwrapper.Errorf("failed to sign trasaction %v, error: %s", tx, err)
		return nil, err
	}

	err = client.SendTransaction(context.TODO(), signedTx)
	if err != nil {
		logwrapper.Error("failed to send trasaction, error: ", err)
		return nil, err
	}
	return signedTx, nil
}
