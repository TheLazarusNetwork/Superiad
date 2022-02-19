package polygon

import (
	"context"
	"math/big"
	"strconv"

	"github.com/TheLazarusNetwork/mtwallet/generated/generc20"
	"github.com/TheLazarusNetwork/mtwallet/generated/generc721"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/envutil"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/wallet"
	rawtrasaction "github.com/TheLazarusNetwork/mtwallet/util/pkg/wallet/rawtransaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
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

func GetBalance(mnemonic string) (*big.Int, error) {
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return nil, err
	}
	publicKey := privKey.PublicKey
	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		return nil, err
	}
	bal, err := client.BalanceAt(context.Background(), crypto.PubkeyToAddress(publicKey), nil)
	if err != nil {
		return nil, err
	}
	return bal, nil
}

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

func GetERC721Balance(mnemonic string, contractAddr common.Address) (*big.Int, error) {
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return nil, err
	}
	publicKey := privKey.PublicKey

	client, err := ethclient.Dial(GetRpcUrl())
	if err != nil {
		return nil, err
	}
	ins, err := generc721.NewErc721(contractAddr, client)
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

func TransferERC721(mnemonic string, toAddress common.Address, contractAddr common.Address, tokenId big.Int) (string, error) {
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
	publicKey := privKey.PublicKey
	fromAddr := crypto.PubkeyToAddress(publicKey)
	tx, err := rawtrasaction.SendRawTrasac(privKey, *client, int64(chainId), 310000, contractAddr, generc721.Erc721MetaData.ABI, "safeTransferFrom", fromAddr, toAddress, &tokenId)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil

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

func GetWalletAddres(mnemonic string) (string, error) {
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return "", err
	}
	walletAddr := crypto.PubkeyToAddress(privKey.PublicKey)

	return walletAddr.String(), nil
}
