package approveall

type ApproveAllRequest struct {
	UserId          string `json:"userId" binding:"required"`
	OperatorAddress string `json:"operatorAddress" binding:"required"`
	ContractAddress string `json:"contractAddress" binding:"required"`
	Approved        bool   `json:"approved" binding:"required"`
}

type TransferPayload struct {
	TrasactionHash string
	Message        string
}
