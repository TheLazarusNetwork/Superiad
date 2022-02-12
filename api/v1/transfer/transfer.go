package transfer

import (
	"math/big"
	"strconv"

	"github.com/TheLazarusNetwork/mtwallet/api/middleware/auth/tokenmiddleware"
	"github.com/TheLazarusNetwork/mtwallet/models/user"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/httphelper"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/network/polygon"
	"github.com/ethereum/go-ethereum/common"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/transfer/:network")
	{
		g.Use(tokenmiddleware.ApiAuth)
		g.GET("/", transfer)
	}
}

func transfer(c *gin.Context) {
	logwrapper.Error(c.Request.URL)
	paramNetwork := c.Param("network")
	erc20Address := c.Query("erc20Address")
	erc721Address := c.Query("erc721Address")
	erc721TokenId := c.Query("erc721TokenId")
	var req TransferRequest
	if err := c.BindJSON(&req); err != nil {
		logwrapper.Errorf("invalid request %v", err.Error())
		httphelper.BadRequest(c)
		return
	}
	mnemonic, err := user.GetMnemonic(req.UserId)
	if err != nil {
		httphelper.
			NewInternalServerError(c,
				"failed to fetch user mnemonic for userId: %v, error: %v",
				req.UserId, err.Error())
		return
	}

	var hash string
	pInfo, err := polygon.GetNetworkInfo()
	if err != nil {
		httphelper.NewInternalServerError(c, "failed to get network info for network %v", "polygon")
		return
	}
	if pInfo.Name == paramNetwork && pInfo.ChainId.Int64() == req.ChainId {
		if len(erc20Address) > 0 {
			hash, err = polygon.TransferERC20(mnemonic, common.HexToAddress(req.To), common.HexToAddress(erc20Address), *big.NewInt(req.Amount))
			if err != nil {
				httphelper.
					NewInternalServerError(c,
						"failed to tranfer to: %v from wallet of userId: %v , network: %v, contractAddr: %v , error: %v", req.To,
						req.UserId, paramNetwork, erc20Address, err.Error())
				return
			}
		} else if len(erc721Address) > 0 {
			if erc721TokenId == "" {
				httphelper.BadRequest(c)
				return
			}
			tokenId, err := strconv.Atoi(erc721TokenId)
			if err != nil {
				httphelper.BadRequest(c)
				return
			}
			hash, err = polygon.TransferERC721(mnemonic, common.HexToAddress(req.To), common.HexToAddress(erc721Address), *big.NewInt(int64(tokenId)))
			if err != nil {
				httphelper.
					NewInternalServerError(c,
						"failed to tranfer to: %v from wallet of userId: %v , network: %v, contractAddr: %v,tokenId %v, error: %v", req.To,
						req.UserId, paramNetwork, erc721Address, tokenId, err.Error())
				return
			}
		} else {
			hash, err = polygon.Transfer(mnemonic, common.HexToAddress(req.To), *big.NewInt(req.Amount))
			if err != nil {
				httphelper.
					NewInternalServerError(c,
						"failed to tranfer to: %v from wallet of userId: %v and network: %v, error: %v", req.To,
						req.UserId, paramNetwork, err.Error())
				return
			}
		}

	} else {
		httphelper.BadRequest(c)
		return
	}

	payload := TransferPayload{
		TrasactionHash: hash,
	}
	httphelper.SuccessResponse(c, "trasaction initiated", payload)
}
