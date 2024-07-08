package repository

import (
	"game/constants"
)

func ExpireOtpTime(key string) bool {

	err := Rdb.Expire(Ctx, key, constants.RateLimitUserExpireTime)
	if err == nil {
		return false
	}

	return true
}
func ExpireWrongedAnswerTime(key string) bool {

	err := Rdb.Expire(Ctx, key, constants.WrongedAnswerExpireTime)
	if err == nil {
		return false
	}

	return true
}
