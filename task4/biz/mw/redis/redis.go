package redis

import (
	"time"

	"github.com/go-redis/redis/v7"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/constants"
)

var (
	expireTime         = time.Hour * 1
	mfaSetupExpireTime = 5 * time.Minute
	rdbPopular         *redis.Client
	rdbMFAAuth         *redis.Client
)

func InitRedis() {
	rdbPopular = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       0,
	})
	rdbMFAAuth = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       1,
	})
}
