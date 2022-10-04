package fetchwallet

import (
	"github.com/TheLazarusNetwork/mtwallet/api/middleware/auth/tokenmiddleware"
	"github.com/TheLazarusNetwork/mtwallet/models/user"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/httphelper"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/network/polygon"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/fetchwallet/:userid")
	{
		g.Use(tokenmiddleware.ApiAuth)
		g.GET("/", fetchwallet)
	}
}

func fetchwallet(c *gin.Context) {
	paramUserId := c.Param("userId")
	if len(paramUserId) <= 0 {
		httphelper.BadRequest(c)
		return
	}
	var req FetchWalletRequest
	err := c.BindJSON(&req)
	if err != nil {
		httphelper.BadRequest(c)
		return
	}
	mnemonic, err := user.GetMnemonic(paramUserId)
	if err != nil {
		httphelper.NewInternalServerError(c, "failed to fetch user mnemonic, error: %v", err.Error())
		return
	}

	walletAddr, err := polygon.GetWalletAddres(mnemonic)
	if err != nil {
		httphelper.NewInternalServerError(c, "failed to get wallet address for userId: %v", paramUserId)
		return
	}
	httphelper.SuccessResponse(c, "wallet address fetched", FetchWalletPayload{
		WalletAddress: walletAddr,
	})

}
