package logic

import (
	"fmt"
	"game/repository"
)

func UpdateNickName(NickName, id string) error {
	primitiveId, err := convertStringToPrimitive(id)
	if err != nil {
		return err
	}
	count, err := repository.FindUserByID(primitiveId)
	if count != 0 {
		return fmt.Errorf("you changed nickname 1 time")
	}
	err = repository.UpdateNickName(primitiveId, NickName)
	if err != nil {
		return err
	}
	return nil

}
