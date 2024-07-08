package logic

import (
	"game/repository"
)

func UpdateNickName(NickName, id string) error {
	primitiveId, err := ConvertStringToPrimitive(id)
	if err != nil {
		return err
	}

	err = repository.UpdateNickName(primitiveId, NickName)
	if err != nil {
		return err
	}
	return nil

}
