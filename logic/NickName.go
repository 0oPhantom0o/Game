package logic

import (
	"game/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NickName(NickName string, primitiveId primitive.ObjectID) error {

	err := repository.UpdateUser(primitiveId, NickName)
	if err != nil {
		return err
	}
	return nil

}
