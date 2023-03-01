package checkbalance

import (
	"github.com/MyriadFlow/superiad/api/v1/matic/checkbalance/checkbalance_erc20"
	"github.com/MyriadFlow/superiad/api/v1/matic/checkbalance/checkbalance_erc721"
	"github.com/MyriadFlow/superiad/api/v1/matic/checkbalance/checkbalance_native"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/checkbalance")
	{

		checkbalance_erc20.ApplyRoutes(g)
		checkbalance_erc721.ApplyRoutes(g)
		checkbalance_native.ApplyRoutes(g)
	}
}
