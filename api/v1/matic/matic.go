package matic

import (
	"github.com/MyriadFlow/superiad/api/v1/matic/approve"
	approveall "github.com/MyriadFlow/superiad/api/v1/matic/approveAll"
	"github.com/MyriadFlow/superiad/api/v1/matic/checkbalance"
	"github.com/MyriadFlow/superiad/api/v1/matic/fetchwallet"
	"github.com/MyriadFlow/superiad/api/v1/matic/isowner"
	signmessage "github.com/MyriadFlow/superiad/api/v1/matic/signMessage"
	"github.com/MyriadFlow/superiad/api/v1/matic/transfer"
	verifysignature "github.com/MyriadFlow/superiad/api/v1/matic/verifySignature"
	"github.com/MyriadFlow/superiad/pkg/paseto"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/matic")
	{
		v1.Use(paseto.PASETO)
		checkbalance.ApplyRoutes(v1)
		fetchwallet.ApplyRoutes(v1)
		isowner.ApplyRoutes(v1)
		verifysignature.ApplyRoutes(v1)

		// v1.Use(onlyunlockedmiddleware.OnlyUnlocked())
		signmessage.ApplyRoutes(v1)
		transfer.ApplyRoutes(v1)
		approve.ApplyRoutes(v1)
		approveall.ApplyRoutes(v1)
	}
}
