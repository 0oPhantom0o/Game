package repository

import (
	"game/constants"
)

func (repo *ConRepository) ExpireOtpTime(key string) bool {

	err := repo.redisdb.Expire(repo.ctx, key, constants.RateLimitUserExpireTime)
	if err == nil {
		return false
	}

	return true
}
func (repo *ConRepository) ExpireWrongedAnswerTime(key string) bool {

	err := repo.redisdb.Expire(repo.ctx, key, constants.WrongedAnswerExpireTime)
	if err == nil {
		return false
	}

	return true
}
