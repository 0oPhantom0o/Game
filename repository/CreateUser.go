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
		fmt.Println(err)
	}
	objectID, Err := result.InsertedID.(primitive.ObjectID)
	if !Err {
		return "", fmt.Errorf("failed to convert inserted ID to ObjectID")
	}

	return objectID.Hex(), nil
}

//a
//func fakerUserAdder(user domain.InternalUser) int {
//	//counter collection documents
//	//	opts2 := options.Count().SetHint("_id_")
//	count, err := collection.CountDocuments(context.TODO(), bson.D{}, opts2)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(count)
//***************************
//make user loop
//counter := 0
//for counter < 50 {
//	user := domain.InternalUser{
//		Phone:         strconv.Itoa(rand.Intn(60)),
//		NickName:      strconv.Itoa(rand.Intn(80)),
//		NickNameLimit: 1,
//		Point:         rand.Intn(100),
//	}
//	counter = fakerUserAdder(user)
//
//	if err != nil {
//		return "", err
//	}
//
//}
//
//collection, err := app.Collection()
//ctx := context.Background()
//
//_, err = collection.InsertOne(ctx, user)
//if err != nil {
//	fmt.Println(err)
//	//}
//	opts2 := options.Count().SetHint("_id_")
//
//	counter, err := collection.CountDocuments(context.TODO(), bson.D{}, opts2)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(counter)
//	return int(counter)
//}
