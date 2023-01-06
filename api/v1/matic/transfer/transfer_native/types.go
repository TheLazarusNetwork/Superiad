package transfer_native

type TransferRequest struct {
	UserId string `json:"userId" binding:"required"`
	To     string `json:"to" binding:"required"`
	Amount int64  `json:"amount" binding:"required"`
}

type TransferPayload struct {
	TrasactionHash string
}

type TransferRequestSalt struct {
	WalletAddress string  `json:"walletAddress" binding:"required"`
	Mnemonic      string  `json:"mnemonic" binding:"required"`
	To            string  `json:"to" binding:"required"`
	Amount        float64 `json:"amount" binding:"required"`
	Salt          string  `json:"salt" binding:"required"`
}
