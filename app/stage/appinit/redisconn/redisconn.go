package redisconn

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/config/envconfig"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/store"
)

func Init() {
	var (
		host     = envconfig.EnvVars.REDIS_HOST
		port     = envconfig.EnvVars.REDIS_PORT
		password = envconfig.EnvVars.REDIS_PASSWORD
	)

	addr := fmt.Sprintf("%s:%d", host, port)

	rds := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       0,        // use default DB
	})

	// Store database in global store
	store.RDS = rds

	// Get underlying sql database to ping it
	pong, err := rds.Ping(context.Background()).Result()

	// If ping fails then log error and exit
	if err != nil {
		logo.Fatal("failed to ping redis", err)
	}

	fmt.Println(pong, err)
}
