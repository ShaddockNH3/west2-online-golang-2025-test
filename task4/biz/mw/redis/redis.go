package redis

import (
	"time"

	"github.com/go-redis/redis/v7"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/constants"
)

var (
	expireTime = time.Hour * 1
	rdbPopular *redis.Client
)

func InitRedis() {
	rdbPopular = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       0,
	})
}
