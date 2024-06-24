package logic

import (
	"game/repository"
)

func NickName(NickName, id string) error {
	primiviteId, err := ConvertStringToPrimivite(id)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	err = repository.UpdateUser(primiviteId, NickName)
	if err != nil {
		return err
	}
	return nil
}
