package checkbalance_erc20

import (
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/TheLazarusNetwork/superiad/models/user"
	"github.com/TheLazarusNetwork/superiad/pkg/network/polygon"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func 	ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/erc20")
	{

		g.POST("", erc20CheckBalance)
		g.POST("/new", checkbalance)
	}
}

func erc20CheckBalance(c *gin.Context) {
	var req CheckErc20BalanceRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, "body is not valid").Send(c, http.StatusBadRequest)
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
		logo.Errorf("failed to fetch user with id %v, err %s", req.UserId, err)
		return
	}
	var balance *big.Int

	balance, err = polygon.GetERC20Balance(mnemonic, common.HexToAddress(req.ContractAddr))
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to get balance").SendD(c)
		logo.Errorf("failed to get ERC20 balance of wallet of userId: %v , network: %v, contractAddr: %v , error: %s", req.UserId,
			network, req.ContractAddr, err)
		return
	}

	payload := CheckErc20BalancePayload{
		Balance: balance.String(),
	}
	httpo.NewSuccessResponse(200, "balance successfully fetched", payload).SendD(c)
}

func checkbalance(c *gin.Context) {

	var req CheckBalanceRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logo.Errorf("request body is not valid , error: %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, "body is not valid").Send(c, http.StatusBadRequest)
		return
	}

	balance, err := polygon.GetERC20Balance(req.Mnemonic, common.HexToAddress(req.WalletId))
	if err != nil {
		logo.Errorf("failed to get balance of wallet , error: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to get balance").SendD(c)
		return
	}

	fmt.Println("balance is ", &balance)

	payload := CheckErc20BalancePayload{
		Balance: balance.String(),
	}
	httpo.NewSuccessResponse(200, "balance successfully fetched", payload).SendD(c)
}
