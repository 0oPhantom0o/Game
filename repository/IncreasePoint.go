package repository

import (
	"fmt"
	"game/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *ConRepository) ChangePoint(id primitive.ObjectID, point int) error {
	collection := repo.mongodb.Database(constants.Database).Collection(constants.UserCollection)

	filter := bson.D{{"_id", id}}
	//update point depends on input -1 or +1
	update := bson.D{{"$inc", bson.D{{"point", point}}}}

	_, err := collection.UpdateOne(repo.ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to increase: %w", err)
	}
	return nil
}
