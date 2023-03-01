package verifysignature

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/MyriadFlow/superiad/models/user"
	"github.com/MyriadFlow/superiad/pkg/network/polygon"
	log "github.com/sirupsen/logrus"
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
		c.JSON(http.StatusBadRequest, err.Error())
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
		}).Error("failed to fetch user mnemonic for userId: %v",
			req.UserId)
		return
	}

	res, err := polygon.VerifySignature(mnemonic, req.Message, req.Signature)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to verify signature")
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to verify signature from wallet of userId: %v and network: %v",
			req.UserId, network)
		return
	}

	if !res {
		c.JSON(401, "signature is invalid")
		return
	}
	c.JSON(200, "signature is valid")
}
