package checkbalance_native

type CheckNativeBalanceRequest struct {
	UserId string `json:"userId" binding:"required"`
}

type CheckNativeBalancePayload struct {
	Balance string `json:"balance"`
}
