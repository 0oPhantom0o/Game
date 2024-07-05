package repository

import (
	"context"
	"game/app"
	"game/constants"
)

func ExpireTime(key string) bool {
	rdb := app.RedisDB
	ctx := context.Background()

	err := rdb.Expire(ctx, key, constants.RateLimitUserExpireTime)
	if err == nil {
		return false
	}

	return true
}
