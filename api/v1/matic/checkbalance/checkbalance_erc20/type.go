package checkbalance_erc20

type CheckErc20BalanceRequest struct {
	UserId       string `json:"userId" binding:"required"`
	ContractAddr string `json:"contractAddr" binding:"required"`
}

type CheckErc20BalancePayload struct {
	Balance string `json:"balance"`
}

type CheckBalanceRequest struct {
	Mnemonic string `json:"mnemonic"`
	WalletId string `json:"walletId"`
}
