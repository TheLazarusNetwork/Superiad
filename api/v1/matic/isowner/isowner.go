package isowner

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
	g := r.Group("/isowner")
	{
		g.Use(tokenmiddleware.ApiAuth)
		g.POST("", isowner)
	}
}

func isowner(c *gin.Context) {
	var req IsOwnerRequest
	err := c.BindJSON(&req)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, "body is invalid").SendD(c)

		return
	}
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

	isOwner, err := polygon.ERC721IsOwner(mnemonic, common.HexToAddress(req.ContractAddress), big.NewInt(req.TokenId))

	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to call ERC721IsOwner").SendD(c)
		logwrapper.Errorf("failed to call ERC721IsOwner for user with id %v,erc721Address %v,tokenId %v, err %s", req.UserId, req.ContractAddress, req.TokenId, err)
		return
	}
	payload := IsOwnerPayload{
		IsOwner: isOwner,
	}
	httpo.NewSuccessResponse(200, "Result success", payload).SendD(c)

}
