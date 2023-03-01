package polygon

import (
	"errors"
	"fmt"

	"github.com/MyriadFlow/superiad/pkg/wallet"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ErrSignatureNotValid = errors.New("signature is not valid")
)

func VerifySignature(mnemonic string, message string, signature string) (bool, error) {
	privKey, err := wallet.GetWallet(mnemonic, GetPath())
	if err != nil {
		return false, err
	}
	walletAddrFromMnemonic := crypto.PubkeyToAddress(privKey.PublicKey)

	newMsg := fmt.Sprintf("\x19Ethereum Signed Message:\n%v%v", len(message), message)
	newMsgHash := crypto.Keccak256Hash([]byte(newMsg))
	signatureInBytes, err := hexutil.Decode(signature)
	if err != nil {
		return false, err
	}
	if signatureInBytes[64] == 27 || signatureInBytes[64] == 28 {
		signatureInBytes[64] -= 27
	}
	pubKeyFromSignature, err := crypto.SigToPub(newMsgHash.Bytes(), signatureInBytes)

	if err != nil {
		return false, err
	}

	//Get address from public key
	walletAddressFromSignature := crypto.PubkeyToAddress(*pubKeyFromSignature)
	if walletAddrFromMnemonic.Hex() == walletAddressFromSignature.Hex() {
		return true, nil
	} else {
		return false, nil
	}

}
