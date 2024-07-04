package logic

import (
	"game/domain"
	"game/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateUser(user *domain.User) (primitive.ObjectID, error) {
	id, err := repository.CreateUser(user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return id, nil
}
