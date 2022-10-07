package setlock

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/TheLazarusNetwork/mtwallet/models/user"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/set-lock")
	{
		g.POST("", setLock)
	}
}

func setLock(c *gin.Context) {
	var req SetLockRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		err := fmt.Errorf("body is invalid: %w", err)
		httpo.NewErrorResponse(http.StatusBadRequest, err.Error()).SendD(c)
		return
	}
	err = user.SetLockStatus(req.Uid, *req.LockStatus)
	if err != nil {
		if errors.Is(user.ErrNoRecordFound, err) {
			httpo.NewErrorResponse(httpo.UserNotFound, "user not found").Send(c, 404)
			return
		}
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to set lock status").SendD(c)
		logo.Errorf("failed to set lock status for user with id: %s, error: %s", req.Uid, err)
		return
	} else if *req.LockStatus {
		httpo.NewSuccessResponse(200, "user locked", nil).SendD(c)
	} else {
		httpo.NewSuccessResponse(200, "user unlocked", nil).SendD(c)
	}
}
