package apiv1

import (
	"github.com/MyriadFlow/superiad/api/v1/auth"
	"github.com/MyriadFlow/superiad/api/v1/matic"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		auth.ApplyRoutes(v1)
		matic.ApplyRoutes(v1)
	}
}
