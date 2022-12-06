package envinit

import "github.com/VirtuaTechnologies/VirtuaCoin_Wallet/config/envconfig"

func Init() {
	envconfig.InitEnvVars()
}
