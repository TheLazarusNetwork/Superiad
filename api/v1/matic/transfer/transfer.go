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
	g := r.Group("/transfer")
	{
		g.Use(tokenmiddleware.ApiAuth)
		g.GET("/", transfer)
	}
}

func transfer(c *gin.Context) {
	network := "matic"
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
	if len(erc20Address) > 0 {
		hash, err = polygon.TransferERC20(mnemonic, common.HexToAddress(req.To), common.HexToAddress(erc20Address), *big.NewInt(req.Amount))
		if err != nil {
			httphelper.
				NewInternalServerError(c,
					"failed to tranfer to: %v from wallet of userId: %v , network: %v, contractAddr: %v , error: %v", req.To,
					req.UserId, network, erc20Address, err.Error())
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
					req.UserId, network, erc721Address, tokenId, err.Error())
			return
		}
	} else {
		hash, err = polygon.Transfer(mnemonic, common.HexToAddress(req.To), *big.NewInt(req.Amount))
		if err != nil {
			httphelper.
				NewInternalServerError(c,
					"failed to tranfer to: %v from wallet of userId: %v and network: %v, error: %v", req.To,
					req.UserId, network, err.Error())
			return
		}
	}
	sendSuccessResponse(c, hash, req.UserId)
	return

}

func sendSuccessResponse(c *gin.Context, hash string, userId string) {
	payload := TransferPayload{
		TrasactionHash: hash,
	}
	if err := user.AddTrasactionHash(userId, hash); err != nil {
		logwrapper.Errorf("failed to add transaction hash: %v to user id: %v, error: %v", hash, userId, err.Error())
	}
	httphelper.SuccessResponse(c, "trasaction initiated", payload)
}
