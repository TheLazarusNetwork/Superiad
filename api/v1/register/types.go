package register

type RegisterPayload struct {
	Uid string `json:"userId"`
}

type WalletResponse struct {
	WalletId string `json:"walletId"`
	Mnemonic string `json:"mnemonic"`
}
