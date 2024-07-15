package repository

import (
	"context"
	"game/app"
	"game/constants"
)

func ExpireOtpTime(key string) bool {
	var rdb = app.RedisDB
	var ctx = context.Background()

	err := rdb.Expire(ctx, key, constants.RateLimitUserExpireTime)
	if err == nil {
		return false
	}

	return true
}
func ExpireWrongedAnswerTime(key string) bool {
	var rdb = app.RedisDB
	var ctx = context.Background()

	err := rdb.Expire(ctx, key, constants.WrongedAnswerExpireTime)
	if err == nil {
		return false
	}

	return true
}

//a
