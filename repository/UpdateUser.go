package repository

import (
	"context"
	"fmt"
	"game/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func UpdateUser(id primitive.ObjectID, phone, nickname string) (string, error) {
	collection, err := app.Collection()
	ctx := context.TODO()
	if err != nil {
		log.Println("Failed to connect to Db : ", err)
		return "", fmt.Errorf("failed to connect to database: %w", err)
	}
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"phone", phone}, {"nickName", nickname}}}}
	result, err := collection.UpdateOne(ctx, filter, update)
	fmt.Println(result)
	return "", nil
}
