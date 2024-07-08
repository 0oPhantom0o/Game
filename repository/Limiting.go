package repository

import (
	"context"
	"game/app"
)

func OTPLimit(phone string) (int64, string) {
	var rdb = app.RedisDB
	var ctx = context.Background()

	//insert rate limit phone and add counter
	rateLimit := "rateLimit : " + phone
	count := rdb.Incr(ctx, rateLimit).Val()

	return count, rateLimit
}
func OTPAnswerLimit(phone string) (int64, string) {
	var rdb = app.RedisDB
	var ctx = context.Background()

	//insert rate limit phone and add counter
	rateLimit := "WrongedAnswerLimit : " + phone
	count := rdb.Incr(ctx, rateLimit).Val()

	return count, rateLimit
}
