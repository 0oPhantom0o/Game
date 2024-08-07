package repository

import (
	"game/constants"
)

func (repo *ConRepository) OtpLimit(phone string) (int64, string) {

	//insert rate limit phone and add counter
	rateLimit := constants.RateLimit + phone
	counter := repo.redisdb.Incr(repo.ctx, rateLimit).Val()

	return counter, rateLimit
}
func (repo *ConRepository) OTPAnswerLimit(phone string) (int64, string) {

	//insert rate limit phone and add counter
	rateLimit := constants.WrongedOtpAnswerLimit + phone
	counter := repo.redisdb.Incr(repo.ctx, rateLimit).Val()

	return counter, rateLimit
}
