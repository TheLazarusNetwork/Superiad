package checkbalance

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
	g := r.Group("/checkbalance/:network")
	{
		g.Use(tokenmiddleware.ApiAuth)
		g.POST("/", checkBalance)
	}
}

func checkBalance(c *gin.Context) {
	var req CheckBalanceRequest
	err := c.BindJSON(&req)
	if err != nil {
		httphelper.BadRequest(c)
		return
	}
	paramNetwork := c.Param("network")
	erc20address := c.Query("erc20address")
	erc721address := c.Query("erc721address")

	mnemonic, err := user.GetMnemonic(req.UserId)
	if err != nil {
		httphelper.NewInternalServerError(c, "failed to fetch user with id %v, err %v", req.UserId, err.Error())
		return
	}
	var balance *big.Int
	pInfo, err := polygon.GetNetworkInfo()
	if err != nil {
		httphelper.NewInternalServerError(c, "failed to get network info for network %v", "polygon")
		return
	}
	if pInfo.Name == paramNetwork && pInfo.ChainId.Int64() == req.ChainId {
		if len(erc20address) > 0 {
			balance, err = polygon.GetERC20Balance(mnemonic, common.HexToAddress(erc20address))
			if err != nil {
				httphelper.
					NewInternalServerError(c,
						"failed to get ERC20 balance of wallet of userId: %v , network: %v, contractAddr: %v , error: %v", req.UserId,
						paramNetwork, erc20address, err.Error())
				return
			}
		} else if len(erc721address) > 0 {
			balance, err = polygon.GetERC721Balance(mnemonic, common.HexToAddress(erc721address))
			if err != nil {
				httphelper.
					NewInternalServerError(c,
						"failed to get ERC721 balance of wallet of userId: %v , network: %v, contractAddr: %v , error: %v", req.UserId,
						paramNetwork, erc721address, err.Error())
				return
			}
		} else {
			balance, err = polygon.GetBalance(mnemonic)
			if err != nil {
				httphelper.
					NewInternalServerError(c,
						"failed to get balance from wallet of userId: %v and network: %v, error: %v",
						req.UserId, paramNetwork, err.Error())
				return
			}
		}
		httphelper.SuccessResponse(c, "balance successfully fetched", CheckBalancePayload{
			Balance: balance.String(),
		})
		return
	}

	httphelper.BadRequest(c)
}
