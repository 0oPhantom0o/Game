package repository

import (
	"context"
	"game/app"
	"game/constants"
)

func OtpLimit(phone string) (int64, string) {
	var rdb = app.RedisDB
	var ctx = context.Background()

	//insert rate limit phone and add counter
	rateLimit := constants.RateLimit + phone
	counter := rdb.Incr(ctx, rateLimit).Val()

	return counter, rateLimit
}
func OTPAnswerLimit(phone string) (int64, string) {
	var rdb = app.RedisDB
	var ctx = context.Background()

	//insert rate limit phone and add counter
	rateLimit := constants.WrongedOtpAnswerLimit + phone
	counter := rdb.Incr(ctx, rateLimit).Val()

	return counter, rateLimit
}
