package auth

type AuthenticatePayload struct {
	Status  int64  `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type AuthenticateReq struct {
	Type string `json:"type" binding:"required"`
}
