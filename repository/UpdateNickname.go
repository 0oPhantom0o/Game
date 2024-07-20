package repository

import (
	"context"
	"fmt"
	"game/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateNickName(id primitive.ObjectID, nickname string) error {
	collection, err := app.Collection()
	ctx := context.TODO()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	filter := bson.D{{"_id", id}}
	//update nickname
	update := bson.D{{"$set", bson.D{{"nickName", nickname}}}, {"$inc", bson.D{{"nick_name_limit", 1}}}}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update nickName: %w", err)
	}
	return nil
}
