package checkbalance_native

import (
	"errors"
	"math/big"
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/TheLazarusNetwork/mtwallet/models/user"
	"github.com/TheLazarusNetwork/mtwallet/pkg/network/polygon"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
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
		httpo.NewErrorResponse(http.StatusBadRequest, "body is invalid").SendD(c)

		return
	}
	network := "matic"

	mnemonic, err := user.GetMnemonic(req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httpo.NewErrorResponse(httpo.UserNotFound, "user not found").Send(c, 404)

			return
		}
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to fetch user").SendD(c)
		logo.Errorf("failed to fetch user with id %v, err %s", req.UserId, err)
		return
	}
	var balance *big.Int

	balance, err = polygon.GetBalance(mnemonic)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to get balance").SendD(c)
		logo.Errorf("failed to get balance from wallet of userId: %v and network: %v, error: %s",
			req.UserId, network, err)
		return
	}

	payload := CheckNativeBalancePayload{
		Balance: balance.String(),
	}
	httpo.NewSuccessResponse(200, "balance successfully fetched", payload).SendD(c)
}
