package config

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func InitRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     AppConfig.Redis.Addr + ":" + AppConfig.Redis.Port,
		Password: AppConfig.Redis.PassWord,
		DB:       AppConfig.Redis.DB,
	})

	ctx := context.Background()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return rdb, nil
}
