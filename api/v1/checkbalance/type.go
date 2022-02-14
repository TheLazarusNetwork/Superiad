package checkbalance

type CheckBalanceRequest struct {
	UserId  string `json:"userId" binding:"required"`
	ChainId int64  `json:"chainId" binding:"required"`
}

type CheckBalancePayload struct {
	Balance string `json:"balance"`
}
