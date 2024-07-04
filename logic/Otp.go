package logic

import (
	"fmt"
	"game/repository"
)

func CheckOtp(phone, code string) error {
	storedCode, err := repository.FindRedisValue(phone)
	if err != nil {
		return err
	}
	if storedCode == code && storedCode != "" {
		return nil
	}
	return fmt.Errorf("error in compare otp")

}
