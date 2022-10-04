package checkbalance

import (
	"github.com/TheLazarusNetwork/mtwallet/api/middleware/auth/tokenmiddleware"
	"github.com/TheLazarusNetwork/mtwallet/api/v1/matic/checkbalance/erc20"
	"github.com/TheLazarusNetwork/mtwallet/api/v1/matic/checkbalance/erc721"
	"github.com/TheLazarusNetwork/mtwallet/api/v1/matic/checkbalance/native"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/checkbalance")
	{
		g.Use(tokenmiddleware.ApiAuth)
		erc20.ApplyRoutes(g)
		erc721.ApplyRoutes(g)
		native.ApplyRoutes(g)
	}
}
