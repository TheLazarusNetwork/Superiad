package verifysignature

type VerifyRequest struct {
	UserId    string `json:"userId" binding:"required"`
	Message   string `json:"message" binding:"required"`
	Signature string `json:"signature" binding:"required,hexadecimal,len=132"`
}
