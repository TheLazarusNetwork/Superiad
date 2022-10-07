package signmessage

import (
	"errors"
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/TheLazarusNetwork/superiad/models/user"
	"github.com/TheLazarusNetwork/superiad/pkg/network/polygon"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/sign-message")
	{

		g.POST("", signMessage)
	}
}

func signMessage(c *gin.Context) {
	var req SignMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logo.Errorf("invalid request %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, "body is invalid").SendD(c)

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

	signature, err := polygon.SignMessage(mnemonic, req.Message)
	if err != nil {

		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to sign").SendD(c)
		logo.Errorf("failed to sign with walletAddress: %s", err)
		return
	}

	sendSuccessResponse(c, signature, req.UserId)

}

func sendSuccessResponse(c *gin.Context, signature string, userId string) {
	payload := SignMessagePayload{
		Signature: signature,
	}
	httpo.NewSuccessResponse(200, "signature generated", payload).SendD(c)
}
