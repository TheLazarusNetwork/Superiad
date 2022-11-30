package transfer_erc20

type TransferRequest struct {
	UserId          string `json:"userId" binding:"required"`
	To              string `json:"to" binding:"required"`
	Amount          int64  `json:"amount" binding:"required"`
	ContractAddress string `json:"contractAddress" binding:"required"`
}

type NewTransferRequest struct {
	UserId       string `json:"userId" binding:"required"`
	RcverWallet  string `json:"rcvwallet" binding:"required"`
	Amount       int64  `json:"amount" binding:"required"`
	SenderWallet string `json:"senderwallet" binding:"required"`
	Mnemonic     string `json:"mnemonic"`
}

type TransferPayload struct {
	TrasactionHash string
}
