package repository

import (
	"context"
	"fmt"
	"game/app"
	"game/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	_, err = collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "phone", Value: 1}, {Key: "nickname", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to insert user into collection because its not uniqe: %w", err)
	}

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to insert user into collection:%w", err)
	}
	mongoID, Err := result.InsertedID.(primitive.ObjectID)
	if !Err {
		return "", fmt.Errorf("failed to convert inserted ID to ObjectID")
	}

	return mongoID.Hex(), nil
}

//
//func fakerUserAdder() int {
//	//counter collection documents
//	collection, err := app.Collection()
//
//	opts2 := options.Count().SetHint("_id_")
//	count, err := collection.CountDocuments(context.TODO(), bson.D{}, opts2)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(count)
//	return int(count)
//}
//
//// ***************************
//// make user loop
//func CreateUser2() {
//	collection, err := app.Collection()
//	ctx := context.Background()
//	if err != nil {
//		fmt.Errorf("failed to connect to database: %w", err)
//	}
//
//	user := domain.InternalUser{
//		Phone:         "",
//		NickName:      "",
//		NickNameLimit: 0,
//		Point:         0,
//	}
//	_, err = collection.Indexes().CreateOne(
//		context.Background(),
//		mongo.IndexModel{
//			Keys:    bson.D{{Key: "phone", Value: 1}, {Key: "nickname", Value: 1}},
//			Options: options.Index().SetUnique(true),
//		},
//	)
//	if err != nil {
//		fmt.Errorf("failed to insert user into collection because its not uniqe: %w", err)
//	}
//	counter := 0
//	for counter < 100 {
//		user := domain.InternalUser{
//			Phone:         strconv.Itoa(rand.Intn(200)),
//			NickName:      strconv.Itoa(rand.Intn(300)),
//			NickNameLimit: 0,
//			Point:         rand.Intn(500),
//		}
//		counter = fakerUserAdder()
//
//		_, err = collection.InsertOne(ctx, user)
//		if err != nil {
//			fmt.Println(err)
//		}
//
//		if err != nil {
//			_ = fmt.Errorf(err.Error())
//		}
//
//	}
//	_, err = collection.InsertOne(ctx, user)
//	if err != nil {
//		_ = fmt.Errorf("failed to insert user into collection:%w", err)
//	}
//}
