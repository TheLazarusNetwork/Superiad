package wallet

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func Test_GetWalletAddr(t *testing.T) {

	mnemonic := "multiply burst shrug insect midnight urban resemble bind fury limit solution diagram scorpion omit feed"
	testc := []struct {
		path         string
		expectedAddr string
	}{
		{
			path:         "m/44H/60H/0H/0/0",
			expectedAddr: "0x4c531Ab9169D4985bB36a94C34496164342D954d",
		},
		{
			path:         "m/44H/60H/0H/0/1",
			expectedAddr: "0x6E71E4e18B901260cB15fC04d52e4BA0846a5584",
		},
	}

	for _, c := range testc {
		privKey, err := GetWallet(mnemonic, c.path)
		if err != nil {
			t.Fatal(err)
		}
		wadr := crypto.PubkeyToAddress(privKey.PublicKey).Hex()
		assert.Equal(t, c.expectedAddr, wadr)
	}
}
