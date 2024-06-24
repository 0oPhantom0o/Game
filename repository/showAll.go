package repository

import (
	"context"
	"fmt"
	"game/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

//	func ShowAllUsers() error {
//		collection, err := app.Collection()
//		ctx := context.TODO()
//		if err != nil {
//			log.Println("Failed to connect to Db : ", err)
//			return fmt.Errorf("failed to connect to database: %w", err)
//		}
//
//		opts := options.Find().SetProjection(bson.D{{"nickName", 1}, {"point", 1}})
//		cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
//		if err != nil {
//			return err
//		}
//		var results []bson.D
//		if err = cursor.All(ctx, &results); err != nil {
//			return err
//		}
//		for _, result := range results {
//			fmt.Println(result)
//		}
//		return nil
//	}
func ShowAllUsers2() error {
	collection, err := app.Collection()
	if err != nil {
		log.Println("Failed to connect to Db : ", err)
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	projectStage := bson.D{{"$project", bson.D{{"nickName", 1}, {"point", 1}}}}
	cursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{projectStage})
	if err != nil {
		return err
	}
	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		return err
	}
	for _, result := range results {
		fmt.Println(result)
	}
	return nil
}
