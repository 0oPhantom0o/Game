package repository

import (
	"context"
	"fmt"
	"game/app"
	"game/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//
//func FindUser(id primitive.ObjectID) (domain.InternalUser, error) {
//	collection, err := app.Collection()
//	var user domain.InternalUser
//	ctx := context.T
//	if err != nil {
//		return user, fmt.Errorf("failed to connect to database: %w", err)
//	}
//	filter := bson.D{{"_id", id}}
//	err = collection.FindOne(ctx, filter).Decode(&user)
//	if err != nil {
//		return user, err
//	}
//	return user, nil
//}

func FindUserByName(name string) (primitive.ObjectID, error) {
	collection, err := app.Collection()
	var user domain.User
	ctx := context.TODO()
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to connect to database: %w", err)
	}
	filter := bson.D{{"phone", name}}

	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return user.ID, nil
}
