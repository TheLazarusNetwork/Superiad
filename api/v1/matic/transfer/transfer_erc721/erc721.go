package transfer_erc721

import (
	"errors"
	"math/big"
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/models/user"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/network/polygon"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/erc721")
	{

		g.POST("", transfer)
	}
}

func transfer(c *gin.Context) {
	network := "matic"
	var req TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logo.Errorf("invalid request %s", err)
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
		logo.Errorf("failed to fetch user mnemonic for userId: %v, error: %s",
			req.UserId, err)
		return
	}

	var hash string
	hash, err = polygon.TransferERC721(mnemonic, common.HexToAddress(req.To), common.HexToAddress(req.ContractAddress), *big.NewInt(int64(req.TokenId)))
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to tranfer").SendD(c)
		logo.Errorf("failed to tranfer to: %v from wallet of userId: %v , network: %v, contractAddr: %v,tokenId %v, error: %s", req.To,
			req.UserId, network, req.ContractAddress, req.TokenId, err)
		return
	}
	sendSuccessResponse(c, hash, req.UserId)
}

func sendSuccessResponse(c *gin.Context, hash string, userId string) {
	payload := TransferPayload{
		TrasactionHash: hash,
	}
	if err := user.AddTrasactionHash(userId, hash); err != nil {
		logo.Errorf("failed to add transaction hash: %v to user id: %v, error: %s", hash, userId, err)
	}
	httpo.NewSuccessResponse(200, "trasaction initiated", payload).SendD(c)
}

func transferWithSalt(c *gin.Context) {
	network := "matic"
	var req TransferRequestSalt
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logo.Errorf("Invalid request body: %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, " Invalid body").SendD(c)
		return
	}
	hash, err := polygon.TransferERC721(req.Mnemonic, common.HexToAddress(req.To), common.HexToAddress(req.ContractAddress), *big.NewInt(int64(req.TokenId)))
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to tranfer").SendD(c)
		logo.Errorf("failed to tranfer to: %v from wallet: %v , network: %v, contractAddr: %v , error: %s", req.To, req.WalletAddress, network, req.ContractAddress, err)
		return
	}

	payload := TransferPayload{TrasactionHash: hash}
	httpo.NewSuccessResponse(http.StatusOK, "trasaction initiated", payload).SendD(c)
}
