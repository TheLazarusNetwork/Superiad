package native

import (
	"math/big"

	"github.com/TheLazarusNetwork/mtwallet/api/middleware/auth/tokenmiddleware"
	"github.com/TheLazarusNetwork/mtwallet/models/user"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/httphelper"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/network/polygon"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/native")
	{
		g.Use(tokenmiddleware.ApiAuth)
		g.POST("/", nativeCheckBalance)
	}
}

func nativeCheckBalance(c *gin.Context) {
	var req CheckNativeBalanceRequest
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

	balance, err = polygon.GetBalance(mnemonic)
	if err != nil {
		httphelper.
			NewInternalServerError(c,
				"failed to get balance from wallet of userId: %v and network: %v, error: %v",
				req.UserId, network, err.Error())
		return
	}
	httphelper.SuccessResponse(c, "balance successfully fetched", CheckNativeBalancePayload{
		Balance: balance.String(),
	})
}
