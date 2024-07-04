package logic

import (
	"fmt"
	"game/repository"
)

func TempUser(phone string) error {
	code, err := randomCode()
	if err != nil {
		return err
	}

	limitCounter, rateLimit := repository.OTPLimit(phone)
	if limitCounter == 1 {
		err := repository.ExpireTime(rateLimit)
		if err != true {
			fmt.Println("error")
		}
	}
	if limitCounter == 5 {
		return fmt.Errorf("user Limited")
	}

	err = repository.RedisDataSet(phone, code, "")
	if err != nil {
		return err
	}
	return nil
}
