package register

type AuthenticatePayload struct {
	Status  int64  `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token"`
}
