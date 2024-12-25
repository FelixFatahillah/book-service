package redis_store

import (
	"book-service/internal/config"
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"
)

func SetKey[T any](ctx context.Context, key string, str T, duration *time.Duration) error {
	redisClient := config.NewRedis()
	var expirationDuration time.Duration

	if duration != nil {
		expirationDuration = *duration
	} else {
		expirationDuration, err := time.ParseDuration(os.Getenv("CACHE_KEY_EXPIRATION"))
		if err != nil {
			log.Printf("Couldn't parse CACHE_KEY_EXPIRATION: %s\n", expirationDuration)
			return err
		}
	}

	err := redisClient.Set(ctx, key, str, expirationDuration).Err()
	if err != nil {
		log.Printf("Couldn't save key: %s, error: %s\n", key, err)
		return err
	}
	return nil
}

func GetKey[T any](ctx context.Context, key string) (*T, error) {
	redisClient := config.NewRedis()
	resultByte, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		if strings.Contains(err.Error(), "redis: nil") {
			return nil, nil
		}
		log.Printf("Couldn't get key: %s, error: %s\n", key, err)
		return nil, err
	}
	var result T
	err = json.Unmarshal([]byte(resultByte), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func DelKey(ctx context.Context, key string) error {
	redisClient := config.NewRedis()
	err := redisClient.Del(ctx, key).Err()
	if err != nil {
		log.Printf("Couldn't del key: %s, error: %s\n", key, err)
		return err
	}
	return nil
}
