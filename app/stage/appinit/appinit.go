package appinit

import (
	"github.com/TheLazarusNetwork/superiad/api/v1/register"
	"github.com/TheLazarusNetwork/superiad/app/stage/appinit/dbconinit"
	"github.com/TheLazarusNetwork/superiad/app/stage/appinit/dbmigrate"
	"github.com/TheLazarusNetwork/superiad/app/stage/appinit/envinit"
	"github.com/TheLazarusNetwork/superiad/app/stage/appinit/loginit"
	"github.com/TheLazarusNetwork/superiad/pkg/auth"
)

func Init() {

	envinit.Init()
	auth.Init()
	register.Init()
	loginit.Init()
	dbconinit.Init()
	dbmigrate.Migrate()
}
