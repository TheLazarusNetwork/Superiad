package checkbalance

import (
	"github.com/TheLazarusNetwork/mtwallet/api/middleware/auth/tokenmiddleware"
	"github.com/TheLazarusNetwork/mtwallet/api/v1/matic/checkbalance/checkbalance_erc20"
	"github.com/TheLazarusNetwork/mtwallet/api/v1/matic/checkbalance/checkbalance_erc721"
	"github.com/TheLazarusNetwork/mtwallet/api/v1/matic/checkbalance/checkbalance_native"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/checkbalance")
	{
		g.Use(tokenmiddleware.ApiAuth)
		checkbalance_erc20.ApplyRoutes(g)
		checkbalance_erc721.ApplyRoutes(g)
		checkbalance_native.ApplyRoutes(g)
	}
}
