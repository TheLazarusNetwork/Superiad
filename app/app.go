package app

import (
	"github.com/TheLazarusNetwork/mtwallet/api"
	"github.com/TheLazarusNetwork/mtwallet/config"
	"github.com/TheLazarusNetwork/mtwallet/models"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/logwrapper"

	"github.com/TheLazarusNetwork/mtwallet/config/dbconfig"

	"github.com/gin-gonic/gin"
)

var GinApp *gin.Engine

func Init(envPath string, logBasePath string) {
	config.Init(envPath)
	logwrapper.Init(logBasePath, true)
	GinApp = gin.Default()
	api.ApplyRoutes(GinApp)
	dbconfig.GetDb()
	models.Migrate()
}
