package appinit

import (
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/app/stage/appinit/envinit"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/app/stage/appinit/logoinit"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/app/stage/appinit/redisconn"
)

func Init() {
	envinit.Init()
	logoinit.Init()
	redisconn.Init()
}
