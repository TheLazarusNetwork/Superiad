package app

import (
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/app/stage/appinit"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/app/stage/apprun"

	"github.com/gin-gonic/gin"
)

var GinApp *gin.Engine

func InitRun() {
	appinit.Init()
	apprun.Run()
}
