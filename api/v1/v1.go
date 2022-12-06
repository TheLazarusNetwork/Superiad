package apiv1

import (
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/matic"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/wallet"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		// register.ApplyRoutes(v1)
		wallet.ApplyRoutes(v1)
		matic.ApplyRoutes(v1)
	}
}
