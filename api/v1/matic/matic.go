package matic

import (
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/matic/checkbalance"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/matic/transfer"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/matic")
	{
		checkbalance.ApplyRoutes(v1)
		// fetchwallet.ApplyRoutes(v1)
		// isowner.ApplyRoutes(v1)
		// verifysignature.ApplyRoutes(v1)

		// signmessage.ApplyRoutes(v1)
		transfer.ApplyRoutes(v1)
		// approve.ApplyRoutes(v1)
		// approveall.ApplyRoutes(v1)
	}
}
