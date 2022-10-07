package fetchwallet

import (
	"errors"
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
	g := r.Group("/fetchwallet/:userId")
	{

		g.GET("", fetchwallet)
	}
}

func fetchwallet(c *gin.Context) {
	paramUserId := c.Param("userId")
	if len(paramUserId) <= 0 {
		httpo.NewErrorResponse(http.StatusBadRequest, "userId is required in /fetchwallet/:userId").SendD(c)
		return
	}

	mnemonic, err := user.GetMnemonic(paramUserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httpo.NewErrorResponse(httpo.UserNotFound, "user not found").Send(c, 404)
			return
		}
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to fetch user").SendD(c)
		logo.Errorf("failed to fetch user mnemonic, error: %s", err)
		return
	}

	walletAddr, err := polygon.GetWalletAddres(mnemonic)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to get wallet address").SendD(c)
		logo.Errorf("failed to get wallet address for userId: %v", paramUserId)
		return
	}

	payload := FetchWalletPayload{
		WalletAddress: walletAddr,
	}
	httpo.NewSuccessResponse(200, "wallet address fetched", payload).SendD(c)

}
