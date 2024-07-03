package logic

import (
	"fmt"
	"game/domain"
	"game/repository"
)

func TempUser(user domain.User) error {
	code, err := randomCode()
	if err != nil {
		return err
	}

	limitCounter, rateLimit := repository.OTPLimit(user.Phone)
	if limitCounter == 1 {
		err := repository.ExpireTime(rateLimit)
		if err != true {
			fmt.Println("error")
		}
	}
	if limitCounter == 5 {
		return fmt.Errorf("user Limited")
	}

	err = repository.TempUser(user.Phone, code)
	if err != nil {
		return err
	}
	return nil
}
