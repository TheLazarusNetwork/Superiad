package checkbalance_native

import (
	"errors"
	"math/big"
	"net/http"

	"github.com/MyriadFlow/superiad/models/user"
	"github.com/MyriadFlow/superiad/pkg/network/polygon"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/native")
	{

		g.POST("", nativeCheckBalance)
	}
}

func nativeCheckBalance(c *gin.Context) {
	var req CheckNativeBalanceRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, "body is invalid")
		return
	}
	network := "matic"

	mnemonic, err := user.GetMnemonic(req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, "user not found")

			return
		}
		c.JSON(http.StatusInternalServerError, "failed to fetch user")
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to fetch user with id %v", req.UserId)
		return
	}
	var balance *big.Int

	balance, err = polygon.GetBalance(mnemonic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to get balance")
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to get balance from wallet of userId: %v and network: %v",
			req.UserId, network)
		return
	}

	payload := CheckNativeBalancePayload{
		Balance: balance.String(),
		Message: "balance successfully fetched",
	}
	c.JSON(200, payload)
}
