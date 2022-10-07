package envinit

import "github.com/TheLazarusNetwork/mtwallet/config/envconfig"

func Init() {
	envconfig.InitEnvVars()
}
