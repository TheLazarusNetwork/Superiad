package appinit

import (
	pasetoauth "github.com/MyriadFlow/superiad/api/v1/auth"
	"github.com/MyriadFlow/superiad/app/stage/appinit/dbconinit"
	"github.com/MyriadFlow/superiad/app/stage/appinit/dbmigrate"
	"github.com/MyriadFlow/superiad/app/stage/appinit/envinit"
	"github.com/MyriadFlow/superiad/app/stage/appinit/loginit"
	"github.com/MyriadFlow/superiad/pkg/auth"
)

func Init() {

	envinit.Init()
	auth.Init()
	pasetoauth.Init()
	loginit.Init()
	dbconinit.Init()
	dbmigrate.Migrate()
}
