package app

import (
	"github.com/TheLazarusNetwork/superiad/app/stage/appinit"
	"github.com/TheLazarusNetwork/superiad/app/stage/apprun"

	"github.com/gin-gonic/gin"
)

var GinApp *gin.Engine

func InitRun() {
	appinit.Init()
	apprun.Run()
}
