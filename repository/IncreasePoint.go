package repository

import (
	"context"
	"fmt"
	"game/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func ChangePoint(point int, id primitive.ObjectID) error {
	collection, err := app.Collection()
	ctx := context.Background()
	if err != nil {
		log.Println("Failed to connect to Db : ", err)
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$inc", bson.D{{"point", point}}}}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to increase: %w", err)
	}
	return nil
}
