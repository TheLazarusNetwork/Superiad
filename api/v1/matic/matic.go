package matic

import (
	"github.com/TheLazarusNetwork/mtwallet/api/v1/matic/approve"
	approveall "github.com/TheLazarusNetwork/mtwallet/api/v1/matic/approveAll"
	"github.com/TheLazarusNetwork/mtwallet/api/v1/matic/checkbalance"
	"github.com/TheLazarusNetwork/mtwallet/api/v1/matic/fetchwallet"
	"github.com/TheLazarusNetwork/mtwallet/api/v1/matic/isowner"
	signmessage "github.com/TheLazarusNetwork/mtwallet/api/v1/matic/signMessage"
	"github.com/TheLazarusNetwork/mtwallet/api/v1/matic/transfer"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/matic")
	{
		transfer.ApplyRoutes(v1)
		checkbalance.ApplyRoutes(v1)
		fetchwallet.ApplyRoutes(v1)
		isowner.ApplyRoutes(v1)
		approve.ApplyRoutes(v1)
		approveall.ApplyRoutes(v1)
		signmessage.ApplyRoutes(v1)
	}
}
