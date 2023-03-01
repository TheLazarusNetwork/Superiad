package isowner

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
	g := r.Group("/isowner")
	{

		g.POST("", isowner)
	}
}

func isowner(c *gin.Context) {
	var req IsOwnerRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
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
		}).Error("failed to fetch user with id %v", req.UserId)
		return
	}

	isOwner, err := polygon.ERC721IsOwner(mnemonic, common.HexToAddress(req.ContractAddress), big.NewInt(req.TokenId))

	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to call ERC721IsOwner")
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to call ERC721IsOwner for user with id %v,erc721Address %v,tokenId %v", req.UserId, req.ContractAddress, req.TokenId)
		return
	}
	payload := IsOwnerPayload{
		IsOwner: isOwner,
		Message: "Result success",
	}
	c.JSON(200, payload)

}
