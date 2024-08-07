package app

import (
	"context"
	"fmt"
	"game/constants"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func mongoConnection() error {
	clientOptions := options.Client().ApplyURI(constants.MongoPort)
	var err error

	MongoClient, err = mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return err
	}
	err = MongoClient.Ping(context.Background(), nil)

	if err != nil {
		return err
	}
	fmt.Println("Connected to MongoDB!")
	return nil
}

func Collection() (*mongo.Client, error) {
	if MongoClient == nil {
		return nil, fmt.Errorf("MongoClient is empty")
	}
	return MongoClient, nil
}
