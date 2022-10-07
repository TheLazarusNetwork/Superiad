package apiv1

import (
	"github.com/TheLazarusNetwork/superiad/api/v1/matic"
	"github.com/TheLazarusNetwork/superiad/api/v1/register"
	"github.com/TheLazarusNetwork/superiad/api/v1/setlock"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		register.ApplyRoutes(v1)
		setlock.ApplyRoutes(v1)
		matic.ApplyRoutes(v1)
	}
}
