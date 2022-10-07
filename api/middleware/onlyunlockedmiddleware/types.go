package onlyunlockedmiddleware

type AccessRequest struct {
	UserId string `json:"userId" binding:"required"`
}
