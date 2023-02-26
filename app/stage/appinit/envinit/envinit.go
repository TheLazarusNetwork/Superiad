package envinit

import (
	"github.com/TheLazarusNetwork/superiad/config/envconfig"
)

func Init() {
	// if os.Getenv("LOAD_CONFIG_FILE") == "" {
	// 	// Load environment variables from .env file
	// 	err := godotenv.Load()
	// 	if err != nil {
	// 		log.WithFields(loginit.StandardFields).Fatalf("Error in reading the config file: %v", err)
	// 	}
	// }
	envconfig.InitEnvVars()
}
