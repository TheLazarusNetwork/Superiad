package isowner

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
	g := r.Group("/isowner/:network")
	{
		g.Use(tokenmiddleware.ApiAuth)
		g.POST("/", isowner)
	}
}

func isowner(c *gin.Context) {
	var req IsOwnerRequest
	err := c.BindJSON(&req)
	if err != nil {
		httphelper.BadRequest(c)
		return
	}
	paramNetwork := c.Param("network")
	erc721Address := c.Query("erc721Address")
	mnemonic, err := user.GetMnemonic(req.UserId)
	if err != nil {
		httphelper.NewInternalServerError(c, "failed to fetch user with id %v, err %v", req.UserId, err.Error())
		return
	}
	pInfo, err := polygon.GetNetworkInfo()
	if err != nil {
		httphelper.NewInternalServerError(c, "failed to get network info for network %v", "polygon")
		return
	}
	if pInfo.Name == paramNetwork && pInfo.ChainId.Int64() == req.ChainId {
		isOwner, err := polygon.ERC721IsOwner(mnemonic, common.HexToAddress(erc721Address), big.NewInt(req.TokenId))

		if err != nil {
			httphelper.NewInternalServerError(c, "failed to run ERC721IsOwner for user with id %v,erc721Address %v,tokenId %v, err %v", req.UserId, erc721Address, req.TokenId, err.Error())
			return
		}
		payload := IsOwnerPayload{
			IsOwner: isOwner,
		}
		httphelper.SuccessResponse(c, "Result success", payload)
	}

}
