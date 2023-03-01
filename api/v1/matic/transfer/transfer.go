package transfer

import (
	"github.com/MyriadFlow/superiad/api/v1/matic/transfer/transfer_erc20"
	"github.com/MyriadFlow/superiad/api/v1/matic/transfer/transfer_erc721"
	"github.com/MyriadFlow/superiad/api/v1/matic/transfer/transfer_native"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/transfer")
	{

		transfer_erc20.ApplyRoutes(g)
		transfer_erc721.ApplyRoutes(g)
		transfer_native.ApplyRoutes(g)
	}
}
