package repository

import (
	"context"
	"fmt"
	"game/app"
	"game/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func CreateUser(user *domain.User) (string, error) {
	collection, err := app.Collection()
	ctx := context.Background()
	if err != nil {
		log.Println("Failed to connect to Db : ", err)
		return "", fmt.Errorf("failed to connect to database: %w", err)
	}

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to insert client into collection: %w", err)
	}
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to convert inserted ID to ObjectID")
	}
	return insertedID.Hex(), nil

}
