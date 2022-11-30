package register

import (
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/TheLazarusNetwork/superiad/models/user"
	"github.com/TheLazarusNetwork/superiad/pkg/network/polygon"
	"github.com/TheLazarusNetwork/superiad/pkg/wallet"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/register")
	{

		g.GET("", register)
		g.GET("/new", newrigister)
	}
}

func register(c *gin.Context) {
	uid, err := user.AddUser()
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to add user").SendD(c)
		logo.Errorf("failed to add user, error: %s", err)

	} else {
		payload := RegisterPayload{
			Uid: uid,
		}
		httpo.NewSuccessResponse(200, "user registration successfull", payload).SendD(c)
	}
}
func newrigister(c *gin.Context) {
	resp, err := createWallet()
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to register ").SendD(c)
		return
	}

	httpo.NewSuccessResponse(200, "user registration successfull", resp).SendD(c)

}

func createWallet() (WalletResponse, error) {

	resp := WalletResponse{}

	mnemonic, err := wallet.GenerateMnemonic()
	if err != nil {
		logo.Errorf("failed to create mnemonic, error: %s", err)
		return WalletResponse{}, err
	}

	walletId, err := polygon.GetWalletAddres(*mnemonic)
	if err != nil {
		logo.Errorf("failed to create wallet, error: %s", err)
		return WalletResponse{}, err
	}

	resp.Mnemonic = *mnemonic
	resp.WalletId = walletId

	return resp, nil

}
