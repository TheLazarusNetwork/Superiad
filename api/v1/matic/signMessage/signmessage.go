package signmessage

import (
	"errors"
	"net/http"

	"github.com/MyriadFlow/superiad/models/user"
	"github.com/MyriadFlow/superiad/pkg/network/polygon"
	log "github.com/sirupsen/logrus"
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
		log.WithFields(log.Fields{
			"err": err,
		}).Error("invalid request")
		c.JSON(http.StatusBadRequest, "body is invalid")

		return
	}
	mnemonic, err := user.GetMnemonic(req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, "user not found")

			return
		}

		c.JSON(http.StatusInternalServerError, "failed to fetch user")
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to fetch user mnemonic for userId: %v", req.UserId)
		return
	}

	signature, err := polygon.SignMessage(mnemonic, req.Message)
	if err != nil {

		c.JSON(http.StatusInternalServerError, "failed to sign")
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to sign with walletAddress")
		return
	}

	sendSuccessResponse(c, signature, req.UserId)

}

func sendSuccessResponse(c *gin.Context, signature string, userId string) {
	payload := SignMessagePayload{
		Signature: signature,
		Message:   "signature generated",
	}
	c.JSON(200, payload)
}
