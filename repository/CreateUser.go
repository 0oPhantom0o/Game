package repository

import (
	"context"
	"fmt"
	"game/app"
	"game/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(phone string) (string, error) {
	collection, err := app.Collection()
	ctx := context.Background()
	if err != nil {
		return "", fmt.Errorf("failed to connect to database: %w", err)
	}

	user := domain.InternalUser{
		Phone:         phone,
		NickName:      "",
		NickNameLimit: 0,
		Point:         0,
	}

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to insert user into collection: %w", err)
	}
	objectID, Err := result.InsertedID.(primitive.ObjectID)
	if !Err {
		return "", fmt.Errorf("failed to convert inserted ID to ObjectID")
	}

	return objectID.Hex(), nil
}

//for i := 0; i < 50; i++ {
//	_, err := collection.InsertOne(ctx, domain.InternalUser{
//		Phone:         strconv.Itoa(rand.Intn(10000)),
//		NickName:      strconv.Itoa(rand.Intn(10000)),
//		NickNameLimit: 1,
//		Point:         rand.Intn(100),
//	})
//	if err != nil {
//		return "", err
//	}
//}
