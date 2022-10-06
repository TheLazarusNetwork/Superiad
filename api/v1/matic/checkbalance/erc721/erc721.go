package erc721

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
	g := r.Group("/erc721")
	{
		g.Use(tokenmiddleware.ApiAuth)
		g.POST("", erc721CheckBalance)
	}
}

func erc721CheckBalance(c *gin.Context) {
	var req CheckErc721BalanceRequest
	err := c.BindJSON(&req)
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
		logwrapper.Errorf("failed to fetch user with id %v, err %s", req.UserId, err)
		return
	}
	var balance *big.Int

	balance, err = polygon.GetERC721Balance(mnemonic, common.HexToAddress(req.ContractAddr))
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to get ERC721 balance").SendD(c)
		logwrapper.Errorf("failed to get ERC721 balance of wallet of userId: %v , network: %v, contractAddr: %v , error: %s", req.UserId,
			network, req.ContractAddr, err)
		return
	}
	httpo.NewSuccessResponse(200, "balance successfully fetched", CheckErc721BalancePayload{
		Balance: balance.String(),
	})
}
