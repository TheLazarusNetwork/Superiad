package signmessage

import (
	"github.com/TheLazarusNetwork/mtwallet/api/middleware/auth/tokenmiddleware"
	"github.com/TheLazarusNetwork/mtwallet/models/user"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/httphelper"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/network/polygon"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/sign-message")
	{
		g.Use(tokenmiddleware.ApiAuth)
		g.GET("/", signMessage)
	}
}

func signMessage(c *gin.Context) {
	var req SignMessageRequest
	if err := c.BindJSON(&req); err != nil {
		logwrapper.Errorf("invalid request %v", err.Error())
		httphelper.BadRequest(c)
		return
	}
	mnemonic, err := user.GetMnemonic(req.UserId)
	if err != nil {
		httphelper.
			NewInternalServerError(c,
				"failed to fetch user mnemonic for userId: %v, error: %v",
				req.UserId, err.Error())
		return
	}

	signature, err := polygon.SignMessage(mnemonic, req.Message)
	if err != nil {
		httphelper.
			NewInternalServerError(c,
				"failed to sign with walletAddress :%s", err.Error())
		return
	}

	sendSuccessResponse(c, signature, req.UserId)
	return

}

func sendSuccessResponse(c *gin.Context, signature string, userId string) {
	payload := SignMessagePayload{
		Signature: signature,
	}
	httphelper.SuccessResponse(c, "signature generated", payload)
}
