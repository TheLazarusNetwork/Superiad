package app

import (
	"github.com/TheLazarusNetwork/mtwallet/api"
	"github.com/TheLazarusNetwork/mtwallet/models"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/logwrapper"

	"github.com/TheLazarusNetwork/mtwallet/config/dbconfig"
	"github.com/TheLazarusNetwork/mtwallet/config/envconfig"

	"github.com/gin-gonic/gin"
)

var GinApp *gin.Engine

func Init() {
	envconfig.InitEnvVars()
	logwrapper.Init()
	GinApp = gin.Default()
	api.ApplyRoutes(GinApp)
	dbconfig.GetDb()
	models.Migrate()
}
