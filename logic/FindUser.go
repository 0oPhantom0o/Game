package logic

import (
	"game/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindUser(name string) (primitive.ObjectID, error) {

	id, err := repository.FindUserByName(name)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return id, nil
}
