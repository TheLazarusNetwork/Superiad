// Package store defines global variables
package store

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RDS *redis.Client
