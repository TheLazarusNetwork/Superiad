package polygon

import (
	"context"
	"math/big"
	"strconv"

	"github.com/TheLazarusNetwork/mtwallet/util/pkg/envutil"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/wallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"golang.org/x/crypto/sha3"
)

func GetChainId() (int, error) {
	chainId, err := strconv.Atoi(envutil.MustGetEnv("NETWORK_CHAIN_ID_POLYGON"))
	if err != nil {
		return 0, err
	}
	return chainId, nil
}

func GetPath() string {
	return envutil.MustGetEnv("NETWORK_PATH_POLYGON")
}

func GetRpcUrl() string {
	return envutil.MustGetEnv("NETWORK_RPC_URL_POLYGON")
}

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
		logwrapper.Errorf("failed to suggestGasTipCap, error %v", err)
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

func TransferERC20(mnemonic string, toAddress common.Address, contractAddr common.Address, amount big.Int) (string, error) {
	tokenAddres := contractAddr
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	// TODO: Ammount, check in goethbook guide
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
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
		logwrapper.Errorf("failed to suggestGasTipCap, error %v", err)
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
		Gas:       310000,
		To:        &tokenAddres,
		Value:     big.NewInt(0),
		Data:      data,
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

func TransferERC721(mnemonic string, toAddress common.Address, contractAddr common.Address, tokenId big.Int) (string, error) {
	tokenAddres := contractAddr
	transferFnSignature := []byte("safeTransferFrom(address,address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return "", err
	}
	publicKey := privKey.PublicKey
	fromAddr := crypto.PubkeyToAddress(publicKey)
	paddedFromAddress := common.LeftPadBytes(fromAddr.Bytes(), 32)
	paddedToAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	paddedTokenId := common.LeftPadBytes(tokenId.Bytes(), 32)
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedFromAddress...)
	data = append(data, paddedToAddress...)
	data = append(data, paddedTokenId...)

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
		logwrapper.Errorf("failed to suggestGasTipCap, error %v", err)
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
		Gas:       310000,
		To:        &tokenAddres,
		Value:     big.NewInt(0),
		Data:      data,
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

func GetNetworkInfo() (*networkInfo, error) {
	chainId, err := GetChainId()
	if err != nil {
		return nil, err
	}
	return &networkInfo{
		Name:    envutil.MustGetEnv("NETWORK_NAME_POLYGON"),
		ChainId: big.NewInt(int64(chainId)),
	}, nil
}
