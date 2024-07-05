package repository

import (
	"context"
	"fmt"
	"game/app"
	"game/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindUserByName(name string) (primitive.ObjectID, error) {
	collection, err := app.Collection()
	var user domain.User
	ctx := context.TODO()
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to connect to database: %w", err)
	}
	filter := bson.D{{"phone", name}}
	//find _id based on phone
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return user.ID, nil
}
