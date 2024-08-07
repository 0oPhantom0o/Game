package repository

import (
	"fmt"
	"game/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *ConRepository) UpdateNickName(id primitive.ObjectID, nickname string) error {
	collection := repo.mongodb.Database(constants.Database).Collection(constants.UserCollection)

	filter := bson.D{{"_id", id}}
	//update nickname
	update := bson.D{{"$set", bson.D{{"nick_name", nickname}}}, {"$inc", bson.D{{"nick_name_limit", 1}}}}
	_, err := collection.UpdateOne(repo.ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update nickName: %w", err)
	}
	return nil
}
