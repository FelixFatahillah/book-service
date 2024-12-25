package config

import (
	"github.com/redis/go-redis/v9"
)

func NewRedis() *redis.Client {
	addr := Viper().GetString("REDIS_ADDR")
	maxPoolSize := Viper().GetInt("REDIS_POOL_MAX_SIZE")
	minIdlePoolSize := Viper().GetInt("REDIS_POOL_MIN_IDLE_SIZE")

	redisStore := redis.NewClient(&redis.Options{
		Addr:         addr,
		PoolSize:     maxPoolSize,
		MinIdleConns: minIdlePoolSize,
	})
	return redisStore
}
