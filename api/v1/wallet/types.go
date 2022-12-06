package wallet

import (
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/models/wallet"
)

type GeneratePayload struct {
	Wallet wallet.Wallet `json:"wallet"`
}
