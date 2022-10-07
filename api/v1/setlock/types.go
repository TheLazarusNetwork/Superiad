package setlock

type SetLockRequest struct {
	LockStatus *bool  `json:"lockStatus" binding:"required"`
	Uid        string `json:"userId" binding:"required"`
}
