package logic

import (
	"fmt"
	"game/domain"
	"game/repository"
)

func CompareOtp(user *domain.User) error {
	storedCode, err := repository.FindRandCode(user.Phone)
	if err != nil {
		return err
	}
	if storedCode == user.RandomCode && storedCode != "" {
		return nil
	}
	return fmt.Errorf("error")

}
