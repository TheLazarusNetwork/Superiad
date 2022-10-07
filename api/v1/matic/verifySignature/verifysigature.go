package verifysignature

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/TheLazarusNetwork/mtwallet/models/user"
	"github.com/TheLazarusNetwork/mtwallet/pkg/network/polygon"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/verify-signature")
	{
		g.POST("", verifySignature)
	}
}

func verifySignature(c *gin.Context) {
	network := "matic"
	var req VerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		err := fmt.Errorf("body is invalid: %w", err)
		httpo.NewErrorResponse(http.StatusBadRequest, err.Error()).SendD(c)
		return
	}
	mnemonic, err := user.GetMnemonic(req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httpo.NewErrorResponse(httpo.UserNotFound, "user not found").Send(c, 404)
			return
		}

		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to fetch user").SendD(c)
		logo.Errorf("failed to fetch user mnemonic for userId: %v, error: %s",
			req.UserId, err)
		return
	}

	res, err := polygon.VerifySignature(mnemonic, req.Message, req.Signature)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to verify signature").SendD(c)
		logo.Errorf("failed to verify signature from wallet of userId: %v and network: %v, error: %s",
			req.UserId, network, err)
		return
	}

	if !res {
		httpo.NewErrorResponse(httpo.SignatureDenied, "signature is invalid").Send(c, 403)
		return
	}
	httpo.NewSuccessResponse(200, "signature is valid", nil).SendD(c)
}
