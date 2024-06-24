package repository

import (
	"context"
	"fmt"
	"game/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func ShowAllUsers() (results []bson.D, err error) {
	collection, err := app.Collection()
	if err != nil {
		log.Println("Failed to connect to Db : ", err)
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	sortStage := bson.D{{"$sort", bson.D{{"point", -1}}}}
	cursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{sortStage})

	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
