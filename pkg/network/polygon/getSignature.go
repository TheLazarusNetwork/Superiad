package polygon

import (
	"fmt"

	"github.com/MyriadFlow/superiad/pkg/wallet"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func SignMessage(mnemonic string, message string) (string, error) {
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return "", err
	}
	newMsg := fmt.Sprintf("\x19Ethereum Signed Message:\n%v%v", len(message), message)

	// keccak256 hash of the data
	dataBytes := []byte(newMsg)
	hashData := crypto.Keccak256Hash(dataBytes)

	signatureBytes, err := crypto.Sign(hashData.Bytes(), privKey)
	if err != nil {
		return "", err
	}

	signature := hexutil.Encode(signatureBytes)

	return signature, nil
}
