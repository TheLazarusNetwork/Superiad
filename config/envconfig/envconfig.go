package envconfig

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	APP_PORT                   int      `env:"APP_PORT,notEmpty"`
	GIN_MODE                   string   `env:"GIN_MODE,notEmpty"`
	DB_HOST                    string   `env:"DB_HOST,notEmpty"`
	DB_USERNAME                string   `env:"DB_USERNAME,notEmpty"`
	DB_PASSWORD                string   `env:"DB_PASSWORD,notEmpty"`
	DB_NAME                    string   `env:"DB_NAME,notEmpty"`
	DB_PORT                    int      `env:"DB_PORT,notEmpty"`
	ALLOWED_ORIGIN             []string `env:"ALLOWED_ORIGIN,notEmpty" envSeparator:","`
	NETWORK_RPC_URL_ETHEREUM   string   `env:"NETWORK_RPC_URL_ETHEREUM,notEmpty"`
	NETWORK_RPC_URL_POLYGON    string   `env:"NETWORK_RPC_URL_POLYGON,notEmpty"`
	TOKEN                      string   `env:"TOKEN,notEmpty"`
	APP_ENVIRONMENT            string   `env:"APP_ENVIRONMENT,notEmpty"`
	PASETO_EXPIRATION_IN_HOURS string   `env:"PASETO_EXPIRATION_IN_HOURS,required" `
	SUPABASEKEY                string   `env:"SUPABASEKEY,required"`
	SUPABASEURL                string   `env:"SUPABASEURL,required"`
	FOOTER                     string   `env:"PASETO_FOOTER,required"`
}

var EnvVars config = config{}

func InitEnvVars() {
	if err := env.Parse(&EnvVars); err != nil {
		log.Fatalf("failed to parse EnvVars: %s", err)
	}
	gin.SetMode(EnvVars.GIN_MODE)
}
