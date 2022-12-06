package wallet

import (
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/models/wallet"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/wallet")
	{
		g.POST("", generateWallet)
		g.GET("/:walletAddress", getWallet)
	}
}

func generateWallet(c *gin.Context) {
	wallet, err := wallet.GenerateWallet()
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to generate wallet").SendD(c)
		logo.Errorf("failed to generate wallet, error: %s", err)

	} else {
		payload := GeneratePayload{
			Wallet: wallet,
		}
		httpo.NewSuccessResponse(200, "wallet creation successful", payload).SendD(c)
	}
}

func getWallet(c *gin.Context) {
	paramWalletAddress := c.Param("walletAddress")
	if len(paramWalletAddress) <= 0 {
		httpo.NewErrorResponse(http.StatusBadRequest, "valid wallet address is required").SendD(c)
		return
	}
	salt, err := wallet.GetWallet(paramWalletAddress)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to get wallet").SendD(c)
		logo.Errorf("failed to get wallet, error: %s", err)
	} else {
		payload := salt
		httpo.NewSuccessResponse(200, "wallet fetched successful ", payload).SendD(c)
	}
}
