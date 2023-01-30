package delegatenftcreation

import (
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1/matic/delegatenftcreation/delegate_erc721"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/delegatenftcreation")
	{
		delegate_erc721.ApplyRoutes(g)
	}
}
