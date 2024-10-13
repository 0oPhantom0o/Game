package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *ConRepository) ChangePoint(id primitive.ObjectID, point int) error {

	filter := bson.D{{"_id", id}}
	//update point depends on input -1 or +1
	update := bson.D{{"$inc", bson.D{{"point", point}}}}

	_, err := repo.mongodb.UpdateOne(repo.ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to increase: %w", err)
	}
	return nil
}
