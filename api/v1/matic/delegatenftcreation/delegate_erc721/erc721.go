package delegate_erc721

import (
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/network/polygon"
	"github.com/ethereum/go-ethereum/common"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/erc721")
	{
		g.POST("", delegateErc721)
	}
}

func delegateErc721(c *gin.Context) {
	network := "matic"
	var req DelegateErc721Request
	if err := c.ShouldBindJSON(&req); err != nil {
		logo.Errorf("invalid request %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, "body is invalid").SendD(c)
		return
	}

	erc721ContractAddr := common.HexToAddress(req.ContractAddress)
	var hash string
	hash, err := polygon.DelegateErc721(req.WalletAddress, erc721ContractAddr, req.MetadataURI)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to tranfer").SendD(c)
		logo.Errorf("failed to delegateAssetCreation of erc721 to wallet Address: %v , network: %v, contractAddr: %v, error: %s", req.WalletAddress, network, req.ContractAddress, err)
		return
	}
	payload := DelegateErc721Payload{
		TrasactionHash: hash,
	}

	httpo.NewSuccessResponse(200, "delegateAssetCreation trasaction initiated", payload).SendD(c)
}
