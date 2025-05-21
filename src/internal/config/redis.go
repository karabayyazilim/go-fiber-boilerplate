package config

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	redisClient *redis.Client
)

var ctx = context.Background()

func Redis() *redis.Client {
	if redisClient != nil {
		return redisClient
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:         Env().RedisHost + ":" + Env().RedisPort,
		DB:           Env().RedisDB,
		Password:     Env().RedisPassword,
		PoolSize:     100,
		MinIdleConns: 20,
		PoolTimeout:  30 * time.Second,
	})

	return redisClient
}
