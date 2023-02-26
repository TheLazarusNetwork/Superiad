package signmessage

type SignMessageRequest struct {
	UserId  string `json:"userId" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type SignMessagePayload struct {
	Signature string `json:"signature"`
	Message string `json:"message"`
}
