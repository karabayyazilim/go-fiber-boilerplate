package redis

import (
	"context"
	"encoding/json"
	"karabayyazilim/src/internal/config"
	"time"
)

func Remember(ctx context.Context, cacheKey string, exprationTime time.Duration, fetchFunc func() (interface{}, error)) (interface{}, error) {
	redisClient := config.Redis()

	cacheData, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var result interface{}
		if err := json.Unmarshal([]byte(cacheData), &result); err == nil {
			return result, nil
		}
	}

	data, err := fetchFunc()
	if err != nil {
		return nil, err
	}

	dataMarshal, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	redisClient.Set(ctx, cacheKey, dataMarshal, exprationTime)
	return data, nil
}
