// Package apprun provides method to Start http server of gin
package apprun

import (
	"fmt"
	"time"

	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/TheLazarusNetwork/mtwallet/api"
	"github.com/TheLazarusNetwork/mtwallet/config/envconfig"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	ginApp := gin.Default()

	corsM := cors.New(cors.Config{AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowOrigins:     envconfig.EnvVars.ALLOWED_ORIGIN})
	ginApp.Use(corsM)
	api.ApplyRoutes(ginApp)
	port := envconfig.EnvVars.APP_PORT
	err := ginApp.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		logo.Fatalf("failed to serve app on port %s: %s", port, err)
	}
}
