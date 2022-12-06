package api

import (
	v1 "github.com/VirtuaTechnologies/VirtuaCoin_Wallet/api/v1"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes Apply the given Routes
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		v1.ApplyRoutes(api)
	}
}
