package fetchwallet

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/MyriadFlow/superiad/models/user"
	"github.com/MyriadFlow/superiad/pkg/network/polygon"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/fetchwallet/:userId")
	{

		g.GET("", fetchwallet)
	}
}

func fetchwallet(c *gin.Context) {
	paramUserId := c.Param("userId")
	if len(paramUserId) <= 0 {
		c.JSON(http.StatusBadRequest, "userId is required in /fetchwallet/:userId")
		return
	}

	mnemonic, err := user.GetMnemonic(paramUserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, "user not found")
			return
		}
		c.JSON(http.StatusInternalServerError, "failed to fetch user")
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to fetch user mnemonic")
		return
	}

	walletAddr, err := polygon.GetWalletAddres(mnemonic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to get wallet address")
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to get wallet address for userId: %v", paramUserId)
		return
	}

	payload := FetchWalletPayload{
		WalletAddress: walletAddr,
		Message:       "wallet address fetched",
	}
	c.JSON(200, payload)

}
