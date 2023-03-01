package approve

import (
	"errors"
	"math/big"
	"net/http"

	"github.com/MyriadFlow/superiad/models/user"
	"github.com/MyriadFlow/superiad/pkg/network/polygon"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/approve")
	{

		g.POST("", approve)
	}
}

func approve(c *gin.Context) {
	network := "matic"
	var req ApproveRequest
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
		}).Error("Failed To fetch user")
		return
	}

	var hash string

	hash, err = polygon.ApproveERC721(mnemonic, common.HexToAddress(req.ToAddress), common.HexToAddress(req.ContractAddress), *big.NewInt(req.TokenId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to approve")
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to approve to: %v from wallet of userId: %v, network: %v, contractAddr: %v, tokenId: %v", req.ToAddress,
			req.UserId, network, req.ContractAddress, req.TokenId)
		return
	}

	SendSuccessResponse(c, hash, req.UserId)
}

func SendSuccessResponse(c *gin.Context, hash string, userId string) {
	payload := TransferPayload{
		TrasactionHash: hash,
		Message:        "trasaction initiated",
	}
	if err := user.AddTrasactionHash(userId, hash); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to add transaction hash: %v to user id: %v", hash, userId)
	}
	c.JSON(200, payload)
}
