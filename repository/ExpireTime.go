package repository

import (
	"context"
	"game/app"
	"time"
)

func ExpireTime(key string) bool {
	rdb := app.RedisDB
	ctx := context.Background()
	err := rdb.Expire(ctx, key, 10*time.Minute)
	if err == nil {
		return false
	}
	return true
}
