package transfer

type TransferRequest struct {
	UserId  string `json:"userId" binding:"required"`
	To      string `json:"to" binding:"required"`
	Amount  int64  `json:"amount" binding:"required"`
	ChainId int64  `json:"chainId" binding:"required"`
}

type TransferPayload struct {
	TrasactionHash string
}
