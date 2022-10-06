package erc20

import (
	"errors"
	"math/big"
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/mtwallet/api/middleware/auth/tokenmiddleware"
	"github.com/TheLazarusNetwork/mtwallet/models/user"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/network/polygon"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/erc20")
	{
		g.Use(tokenmiddleware.ApiAuth)
		g.POST("", erc20CheckBalance)
	}
}

func erc20CheckBalance(c *gin.Context) {
	var req CheckErc20BalanceRequest
	err := c.BindJSON(&req)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, "body is not valid").Send(c, http.StatusBadRequest)
		return
	}
	network := "matic"

	mnemonic, err := user.GetMnemonic(req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httpo.NewErrorResponse(httpo.UserNotFound, "user not found").SendD(c)
			return
		}
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to fetch user").SendD(c)
		logwrapper.Errorf("failed to fetch user with id %v, err %s", req.UserId, err)
		return
	}
	var balance *big.Int

	balance, err = polygon.GetERC20Balance(mnemonic, common.HexToAddress(req.ContractAddr))
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to get balance").SendD(c)
		logwrapper.Errorf("failed to get ERC20 balance of wallet of userId: %v , network: %v, contractAddr: %v , error: %s", req.UserId,
			network, req.ContractAddr, err)
		return
	}

	httpo.NewSuccessResponse(200, "balance successfully fetched", CheckErc20BalancePayload{
		Balance: balance.String(),
	}).SendD(c)
}
