package logic

import (
	"fmt"
	"game/repository"
)

func UpdateNickName(nickName, id string) error {
	primitiveId, err := convertStringToPrimitive(id)
	if err != nil {
		return err
	}
	count, err := repository.FindUserByID(primitiveId)
	if count >= 1 {
		return fmt.Errorf("you changed nickname 2 times")
	}
	err = repository.UpdateNickName(primitiveId, nickName)
	if err != nil {
		return err
	}
	return nil

}
