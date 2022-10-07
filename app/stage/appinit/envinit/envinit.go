package envinit

import "github.com/TheLazarusNetwork/superiad/config/envconfig"

func Init() {
	envconfig.InitEnvVars()
}
