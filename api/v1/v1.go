package apiv1

import (
	"github.com/TheLazarusNetwork/mtwallet/api/v1/register"
	"github.com/TheLazarusNetwork/mtwallet/api/v1/transfer"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		register.ApplyRoutes(v1)
		transfer.ApplyRoutes(v1)
	}
}
