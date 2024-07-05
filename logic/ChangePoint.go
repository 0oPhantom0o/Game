package logic

import (
	"game/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ChangePoint(id primitive.ObjectID, point int) error {

	err := repository.ChangePoint(id, point)
	if err != nil {
		return err
	}
	return nil
}
