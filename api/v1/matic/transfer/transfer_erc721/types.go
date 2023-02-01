package transfer_erc721

type TransferRequest struct {
	UserId          string `json:"userId" binding:"required"`
	To              string `json:"to" binding:"required"`
	Amount          int64  `json:"amount" binding:"required"`
	ContractAddress string `json:"contractAddress" binding:"required"`
	TokenId         int64  `json:"tokenId" binding:"required"`
}

type TransferPayload struct {
	TrasactionHash string
}

type TransferRequestSalt struct {
	WalletAddress   string  `json:"walletAddress" binding:"required"`
	Mnemonic        string  `json:"mnemonic" binding:"required"`
	To              string  `json:"to" binding:"required"`
	TokenId         float64 `json:"amount" binding:"required"`
	ContractAddress string  `json:"contractAddress" binding:"required"`
	Salt            string  `json:"salt" binding:"required"`
}
