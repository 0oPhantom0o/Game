package logic

import (
	"game/repository"
)

func checkOtp(phone, code string) (bool, error) {
	storedCode, err := repository.FindStoredOtp(phone)
	if err != nil {
		return false, err
	}
	if storedCode == code && storedCode != "" {
		return true, nil
	}

	return false, nil

}
