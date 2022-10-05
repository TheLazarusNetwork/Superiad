package approve

import (
	"math/big"

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
	g := r.Group("/approve")
	{
		g.Use(tokenmiddleware.ApiAuth)
		g.GET("/", approve)
	}
}

func approve(c *gin.Context) {
	network := "matic"
	var req ApproveRequest
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

	hash, err = polygon.ApproveERC721(mnemonic, common.HexToAddress(req.ToAddress), common.HexToAddress(req.ContractAddress), *big.NewInt(req.TokenId))
	if err != nil {
		httphelper.
			NewInternalServerError(c,
				"failed to approve to: %v from wallet of userId: %v, network: %v, contractAddr: %v, tokenId: %v, error: %v", req.ToAddress,
				req.UserId, network, req.ContractAddress, req.TokenId, err.Error())
		return
	}

	sendSuccessResponse(c, hash, req.UserId)
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
