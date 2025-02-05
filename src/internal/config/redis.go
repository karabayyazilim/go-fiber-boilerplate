package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
)

var ctx = context.Background()

func Redis() *redis.Client {
	redisClient = redis.NewClient(&redis.Options{
		Addr:           Env().RedisHost + ":" + Env().RedisPort,
		DB:             Env().RedisDB,
		Password:       Env().RedisPassword,
		MaxActiveConns: 30,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Redis Error: ", err)
	}

	return redisClient
}
