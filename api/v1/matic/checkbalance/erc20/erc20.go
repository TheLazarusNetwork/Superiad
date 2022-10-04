package erc20

import (
	"math/big"

	"github.com/TheLazarusNetwork/mtwallet/api/middleware/auth/tokenmiddleware"
	"github.com/TheLazarusNetwork/mtwallet/models/user"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/httphelper"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/network/polygon"
	"github.com/ethereum/go-ethereum/common"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/erc20")
	{
		g.Use(tokenmiddleware.ApiAuth)
		g.POST("/", erc20CheckBalance)
	}
}

func erc20CheckBalance(c *gin.Context) {
	var req CheckErc20BalanceRequest
	err := c.BindJSON(&req)
	if err != nil {
		httphelper.BadRequest(c)
		return
	}
	network := "matic"

	mnemonic, err := user.GetMnemonic(req.UserId)
	if err != nil {
		httphelper.NewInternalServerError(c, "failed to fetch user with id %v, err %v", req.UserId, err.Error())
		return
	}
	var balance *big.Int

	balance, err = polygon.GetERC20Balance(mnemonic, common.HexToAddress(req.ContractAddr))
	if err != nil {
		httphelper.
			NewInternalServerError(c,
				"failed to get ERC20 balance of wallet of userId: %v , network: %v, contractAddr: %v , error: %v", req.UserId,
				network, req.ContractAddr, err.Error())
		return
	}
	httphelper.SuccessResponse(c, "balance successfully fetched", CheckErc20BalancePayload{
		Balance: balance.String(),
	})
	return

}
