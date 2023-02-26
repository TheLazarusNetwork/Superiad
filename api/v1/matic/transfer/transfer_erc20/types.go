package transfer_erc20

type TransferRequest struct {
	UserId          string `json:"userId" binding:"required"`
	To              string `json:"to" binding:"required"`
	Amount          int64  `json:"amount" binding:"required"`
	ContractAddress string `json:"contractAddress" binding:"required"`
}

type TransferPayload struct {
	TrasactionHash string
	Message        string `json:"message"`
}
