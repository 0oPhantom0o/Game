package repository

import (
	"context"
	"fmt"
	"game/app"
	"game/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func FindUser(id primitive.ObjectID) (domain.InternalUser, error) {
	collection, err := app.Collection()
	var user domain.InternalUser
	ctx := context.TODO()
	if err != nil {
		log.Println("Failed to connect to Db : ", err)
		return user, fmt.Errorf("failed to connect to database: %w", err)
	}
	filter := bson.D{{"_id", id}}
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
