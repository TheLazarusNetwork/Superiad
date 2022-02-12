package api

import (
	"github.com/TheLazarusNetwork/mtwallet/api/middleware/auth/tokenmiddleware"
	v1 "github.com/TheLazarusNetwork/mtwallet/api/v1"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes Apply the given Routes
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.Use(tokenmiddleware.ApiAuth)
		v1.ApplyRoutes(api)
	}
}
