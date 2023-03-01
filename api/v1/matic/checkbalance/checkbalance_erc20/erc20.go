package checkbalance_erc20

import (
	"errors"
	"math/big"
	"net/http"

	"github.com/MyriadFlow/superiad/models/user"
	"github.com/MyriadFlow/superiad/pkg/network/polygon"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/erc20")
	{

		g.POST("", erc20CheckBalance)
	}
}

func erc20CheckBalance(c *gin.Context) {
	var req CheckErc20BalanceRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, "body is not valid")
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

	balance, err = polygon.GetERC20Balance(mnemonic, common.HexToAddress(req.ContractAddr))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to get balance")
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to get ERC20 balance of wallet of userId: %v , network: %v, contractAddr: %v ", req.UserId,
			network, req.ContractAddr)
		return
	}

	payload := CheckErc20BalancePayload{
		Balance: balance.String(),
		Message: "balance successfully fetched",
	}
	c.JSON(200, payload)
}
