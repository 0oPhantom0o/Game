package logic

import (
	"game/domain"
	"game/repository"
)

func TempUser(user domain.User) error {
	code, err := randomCode()
	if err != nil {
		return err
	}
	err = repository.TempUser(user.Phone, code)
	if err != nil {
		return err
	}
	return nil
}
