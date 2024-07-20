package repository

import (
	"context"
	"game/app"
)

func OtpLimit(phone string) (int64, string) {
	var rdb = app.RedisDB
	var ctx = context.Background()

	//insert rate limit phone and add counter
	rateLimit := "rateLimit:" + phone
	limitCounter := rdb.Incr(ctx, rateLimit).Val()

	return limitCounter, rateLimit
}
func OTPAnswerLimit(phone string) (int64, string) {
	var rdb = app.RedisDB
	var ctx = context.Background()

	//insert rate limit phone and add counter
	rateLimit := "WrongedOTPAnswerLimit:" + phone
	limitCounter := rdb.Incr(ctx, rateLimit).Val()

	return limitCounter, rateLimit
}

//
//func AnswerLimit(id string) (int64, string) {
//	var rdb = app.RedisDB
//	var ctx = context.Background()
//
//	//insert rate limit phone and add counter
//	rateLimit := "QuestionAnswerLimit : " + id
//	count := rdb.Incr(ctx, rateLimit).Val()
//
//	return count, rateLimit
//}
//a
