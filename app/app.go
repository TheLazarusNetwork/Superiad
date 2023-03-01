package app

import (
	"github.com/MyriadFlow/superiad/app/stage/appinit"
	"github.com/MyriadFlow/superiad/app/stage/apprun"

	"github.com/gin-gonic/gin"
)

var GinApp *gin.Engine

func InitRun() {
	appinit.Init()
	apprun.Run()
}
