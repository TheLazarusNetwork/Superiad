package approveall

import (
	"errors"
	"net/http"

	"github.com/MyriadFlow/superiad/api/v1/matic/approve"
	"github.com/MyriadFlow/superiad/models/user"
	"github.com/MyriadFlow/superiad/pkg/network/polygon"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/approve-all")
	{
		g.POST("", approveAll)
	}
}

func approveAll(c *gin.Context) {
	network := "matic"
	var req ApproveAllRequest
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

	hash, err = polygon.SetAprovalForAllErc721(mnemonic, common.HexToAddress(req.OperatorAddress), common.HexToAddress(req.ContractAddress), req.Approved)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to approve")
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to approve all to operator: %v from wallet of userId: %v, network: %v, contractAddr: %v", req.OperatorAddress, req.UserId, network, req.ContractAddress)
		return
	}

	approve.SendSuccessResponse(c, hash, req.UserId)
}

// func sendSuccessResponse(c *gin.Context, hash string, userId string) {
// 	payload := TransferPayload{
// 		TrasactionHash: hash,
// 		Message:        "trasaction initiated",
// 	}
// 	if err := user.AddTrasactionHash(userId, hash); err != nil {
// 		log.WithFields(log.Fields{
// 			"err": err,
// 		}).Error("failed to add transaction hash: %v to user id: %v, error: %s", hash, userId, err)
// 	}
// 	c.JSON(200, payload)
// }
