package app

import (
	"github.com/TheLazarusNetwork/mtwallet/app/stage/appinit"
	"github.com/TheLazarusNetwork/mtwallet/app/stage/apprun"

	"github.com/gin-gonic/gin"
)

var GinApp *gin.Engine

func InitRun() {
	appinit.Init()
	apprun.Run()
}
