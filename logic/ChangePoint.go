package logic

import (
	"game/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ChangePoint(id primitive.ObjectID, point int) error {

	//userData, err := repository.FindUser(id)
	//if err != nil {
	//	return err
	//}
	//userData.Point = point
	err := repository.ChangePoint(point, id)
	if err != nil {
		return err
	}
	return nil
}
