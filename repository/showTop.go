package repository

import (
	"context"
	"fmt"
	"game/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ShowAllUsers(number int64) (results []bson.D, err error) {

	collection, err := app.Collection()
	if err != nil {
		log.Println("Failed to connect to Db : ", err)
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"point", -1}}).SetLimit(number)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
