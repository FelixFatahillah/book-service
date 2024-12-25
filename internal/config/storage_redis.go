package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/redis/v3"
	"runtime"
)

func NewStorageRedis() fiber.Storage {
	return redis.New(redis.Config{
		Addrs:    []string{Viper().GetString("REDIS_ADDR")},
		Password: Viper().GetString("REDIS_PASSWORD"),
		Database: Viper().GetInt("REDIS_DB"),
		PoolSize: 10 * runtime.GOMAXPROCS(0),
	})
}
