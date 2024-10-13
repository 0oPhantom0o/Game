package logic

import (
	"fmt"
)

func (g *GameLogic) UpdateNickName(nickName, id string) error {
	primitiveId, err := convertStringToPrimitive(id)
	if err != nil {
		return err
	}
	count, err := g.Repo.FindUserByID(primitiveId)
	if count >= 1 {
		return fmt.Errorf("you changed nickname before")
	}
	err = g.Repo.UpdateNickName(primitiveId, nickName)
	if err != nil {
		return err
	}
	return nil

}
