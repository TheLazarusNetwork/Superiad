package transfer_native

type TransferRequest struct {
	UserId string `json:"userId" binding:"required"`
	To     string `json:"to" binding:"required"`
	Amount int64  `json:"amount" binding:"required"`
}

