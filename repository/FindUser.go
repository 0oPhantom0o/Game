package repository

import (
	"context"
	"fmt"
	"game/app"
	"game/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func FindUserIdByPhone(phone string) (string, error) {
	collection, err := app.Collection()
	var user domain.UserId
	ctx := context.TODO()
	if err != nil {
		return "", fmt.Errorf("failed to connect to database: %w", err)
	}
	filter := bson.D{{"phone", phone}}
	//find _id based on phone
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return "", err
	}
	return user.ID.Hex(), nil
}
