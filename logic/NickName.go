package logic

import (
	"game/domain"
	"game/repository"
)

func NickName(user *domain.User) error {
	id := user.ID
	nickname := user.NickName
	phone, err := repository.FindUser(id)
	if err != nil {
		return err
	}
	nickname, err = repository.UpdateUser(id, phone.Phone, nickname)
	if err != nil {
		return err
	}
	return nil
}
