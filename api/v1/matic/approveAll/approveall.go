package approveall

import (
	"errors"
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/TheLazarusNetwork/mtwallet/models/user"
	"github.com/TheLazarusNetwork/mtwallet/pkg/network/polygon"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/approve-all")
	{

		g.POST("", approveAll)
	}
}

func approveAll(c *gin.Context) {
	network := "matic"
	var req ApproveAllRequest
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

	hash, err = polygon.SetAprovalForAllErc721(mnemonic, common.HexToAddress(req.OperatorAddress), common.HexToAddress(req.ContractAddress), req.Approved)
	if err != nil {

		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to approve all").SendD(c)
		logo.Errorf("failed to approve all to operator: %v from wallet of userId: %v, network: %v, contractAddr: %v, error: %s", req.OperatorAddress,
			req.UserId, network, req.ContractAddress, err)
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
