package repository

import (
	"context"
	"fmt"
	"game/app"
	"game/domain"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func IncreasePoint(user domain.Internaluser) error {
	collection, err := app.Collection()
	ctx := context.Background()
	if err != nil {
		log.Println("Failed to connect to Db : ", err)
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	filter := bson.D{{"_id", user.ID}}
	update := bson.D{{"$set", bson.D{{"phone", user.Phone}, {"nickName", user.NickName}}}, {"$inc", bson.D{{"point", 1}}}}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to increase: %w", err)
	}
	return nil
}
