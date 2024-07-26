package app

import (
	"context"
	"fmt"
	"game/constants"
	"github.com/redis/go-redis/v9"
)

var RedisDB *redis.Client

func redisConnection() error {
	ctx := context.Background()
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     constants.RedisPort,
		Password: "",
		DB:       0,
	})
	_, err := RedisDB.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("error initializing Redis database:%w", err)
	}

	return nil

}
