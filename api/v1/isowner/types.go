package isowner

type IsOwnerRequest struct {
	UserId  string `json:"userId" binding:"required"`
	ChainId int64  `json:"chainId" binding:"required"`
	TokenId int64  `json:"tokenId" binding:"required"`
}

type IsOwnerPayload struct {
	IsOwner bool `json:"isOwner" binding:"required"`
}
