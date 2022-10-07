package appinit

import (
	"github.com/TheLazarusNetwork/mtwallet/app/stage/appinit/dbconinit"
	"github.com/TheLazarusNetwork/mtwallet/app/stage/appinit/dbmigrate"
	"github.com/TheLazarusNetwork/mtwallet/app/stage/appinit/envinit"
	"github.com/TheLazarusNetwork/mtwallet/app/stage/appinit/logoinit"
)

func Init() {
	envinit.Init()
	logoinit.Init()
	dbconinit.Init()
	dbmigrate.Migrate()
}
