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
	opts := options.Find().SetSort(bson.D{{"point", -1}}).SetLimit(number).SetProjection(bson.D{{"phone", 0}, {"_id", 0}})
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}
	fmt.Println(results)
	return results, nil
}
