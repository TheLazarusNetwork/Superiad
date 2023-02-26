package isowner

type IsOwnerRequest struct {
	UserId          string `json:"userId" binding:"required"`
	TokenId         int64  `json:"tokenId" binding:"required"`
	ContractAddress string `json:"contractAddress" binding:"required"`
}

type IsOwnerPayload struct {
	IsOwner bool `json:"isOwner" binding:"required"`
	Message string `json:"message"`
}
