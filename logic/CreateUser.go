package logic

import (
	"game/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateUser(phone string) (primitive.ObjectID, error) {
	id, err := repository.CreateUser(phone)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return id, nil
}
