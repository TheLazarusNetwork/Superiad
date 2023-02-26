package checkbalance_erc721

type CheckErc721BalanceRequest struct {
	UserId       string `json:"userId" binding:"required"`
	ContractAddr string `json:"contractAddr" binding:"required"`
}

type CheckErc721BalancePayload struct {
	Balance string `json:"balance"`
	Message string `json"message"`
}
