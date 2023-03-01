package transfer_native

import (
	"errors"
	"math/big"
	"net/http"

	"github.com/MyriadFlow/superiad/api/v1/matic/transfer/transfer_erc20"
	"github.com/MyriadFlow/superiad/models/user"
	"github.com/MyriadFlow/superiad/pkg/network/polygon"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/native")
	{
		g.POST("", nativeTransfer)
	}
}

func nativeTransfer(c *gin.Context) {
	network := "matic"
	var req TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Invalid Request")
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

	var hash string
	hash, err = polygon.Transfer(mnemonic, common.HexToAddress(req.To), *big.NewInt(req.Amount))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to tranfer")
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to tranfer to: %v from wallet of userId: %v and network: %v", req.To,
			req.UserId, network)
		return
	}
	transfer_erc20.SendSuccessResponse(c, hash, req.UserId)
}
