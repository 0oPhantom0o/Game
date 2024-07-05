package logic

import (
	"fmt"
	"game/repository"
)

func TempUser(phone, code string) error {
	//limiting request otp
	limitCounter, rateLimit := repository.OTPLimit(phone)
	if limitCounter == 1 {
		err := repository.ExpireTime(rateLimit)
		if err != true {
			return fmt.Errorf("error in expire time set")
		}
	}
	if limitCounter == 5 {
		return fmt.Errorf("user Limited")
	}

	err := repository.RedisDataSet(phone, code, "TempUser")
	if err != nil {
		return fmt.Errorf("couldnt insert user in temp database")
	}
	return nil
}
