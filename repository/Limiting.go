package repository

import (
	"context"
	"game/app"
)

func OTPLimit(phone string) (int64, string) {
	rdb := app.RedisDB
	ctx := context.Background()
	rateLimit := "rateLimit : " + phone
	count := rdb.Incr(ctx, rateLimit).Val()

	return count, rateLimit
}
