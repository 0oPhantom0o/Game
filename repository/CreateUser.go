package repository

import (
	"context"
	"fmt"
	"game/app"
	"game/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(client *domain.User) (primitive.ObjectID, error) {
	collection, err := app.Collection()
	ctx := context.Background()
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to connect to database: %w", err)
	}
	user := domain.InternalUser{
		Phone:    client.Phone,
		NickName: "",
		Point:    0,
	}
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to insert user into collection: %w", err)
	}
	objectID, Err := result.InsertedID.(primitive.ObjectID)
	if !Err {
		return primitive.NilObjectID, fmt.Errorf("failed to convert inserted ID to ObjectID")
	}
	return objectID, nil

}

//user2 := domain.InternalUser{
//	Phone:    "akbar",
//	NickName: "",
//	Point:    5,
//}
//user3 := domain.InternalUser{
//	Phone:    "ali",
//	NickName: "",
//	Point:    3,
//}
//user4 := domain.InternalUser{
//	Phone:    "asqar",
//	NickName: "",
//	Point:    15,
//}
//user5 := domain.InternalUser{
//	Phone:    "mohsen",
//	NickName: "",
//	Point:    2,
//}
//user6 := domain.InternalUser{
//	Phone:    "client.Phone",
//	NickName: "",
//	Point:    -2,
//}
//_, _ = collection.InsertOne(ctx, user2)
//_, _ = collection.InsertOne(ctx, user3)
//_, _ = collection.InsertOne(ctx, user4)
//_, _ = collection.InsertOne(ctx, user5)
//_, _ = collection.InsertOne(ctx, user6)
