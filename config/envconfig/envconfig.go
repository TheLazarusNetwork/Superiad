package envconfig

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	APP_PORT                 int      `env:"APP_PORT,notEmpty"`
	GIN_MODE                 string   `env:"GIN_MODE,notEmpty"`
	REDIS_HOST               string   `env:"REDIS_HOST,notEmpty"`
	REDIS_PORT               int      `env:"REDIS_PORT,notEmpty"`
	REDIS_PASSWORD           string   `env:"REDIS_PASSWORD,notEmpty"`
	REDIS_DB                 string   `env:"REDIS_DB,notEmpty"`
	ALLOWED_ORIGIN           []string `env:"ALLOWED_ORIGIN,notEmpty" envSeparator:","`
	NETWORK_RPC_URL_ETHEREUM string   `env:"NETWORK_RPC_URL_ETHEREUM,notEmpty"`
	NETWORK_RPC_URL_POLYGON  string   `env:"NETWORK_RPC_URL_POLYGON,notEmpty"`
	TOKEN                    string   `env:"TOKEN,notEmpty"`
	APP_ENVIRONMENT          string   `env:"APP_ENVIRONMENT,notEmpty"`
}

var EnvVars config = config{}

func InitEnvVars() {
	if err := env.Parse(&EnvVars); err != nil {
		log.Fatalf("failed to parse EnvVars: %s", err)
	}
	gin.SetMode(EnvVars.GIN_MODE)
}
