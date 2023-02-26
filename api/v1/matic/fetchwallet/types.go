package fetchwallet

type FetchWalletPayload struct {
	WalletAddress string `json:"walletAddress"`
	Message       string `json:"message"`
}
