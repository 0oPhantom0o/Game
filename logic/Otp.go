package logic

import (
	"fmt"
	"game/repository"
)

func CompareOtp(phone, code string) error {
	storedCode, err := repository.FindUserTempData(phone)
	if err != nil {
		return err
	}
	if storedCode == code && storedCode != "" {
		return nil
	}
	return fmt.Errorf("error in compare otp")

}
