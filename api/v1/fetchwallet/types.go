package fetchwallet

type FetchWalletRequest struct {
	Network string `json:"network" binding:"required"`
	ChainId int64  `json:"chainId" binding:"required"`
}
type FetchWalletPayload struct {
	WalletAddress string `json:"walletAddress"`
}
