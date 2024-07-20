package logic

import (
	"game/repository"
)

func checkOtp(phone, randomCode string) (bool, error) {
	storedCode, err := repository.FindStoredOtp(phone)
	if err != nil {
		return false, err
	}
	if storedCode == randomCode && storedCode != "" {
		return true, nil
	}

	return false, nil

}
