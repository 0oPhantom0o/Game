package repository

import (
	"context"
	"fmt"
	"game/app"
	"game/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func FindUserByID(id primitive.ObjectID) (int, error) {
	collection, err := app.Collection()
	var user domain.InternalUser
	ctx := context.TODO()
	if err != nil {
		return 0, fmt.Errorf("failed to connect to database: %w", err)
	}
	filter := bson.D{{"_id", id}}
	//find _id based on phone
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return 0, err
	}
	return user.NickNameLimit, nil
}

func FindStoredOtp(id string) (string, error) {
	var rdb = app.RedisDB
	var ctx = context.Background()

	value, err := rdb.Get(ctx, id).Result()
	if err != nil {
		return "", fmt.Errorf("user doesnt exist")
	}
	return value, nil

}
func FindAnswer(id string) (string, error) {
	var rdb = app.RedisDB
	var ctx = context.Background()

	value, err := rdb.Get(ctx, id).Result()
	if err != nil {
		return "", err
	}
	return value, nil

}
