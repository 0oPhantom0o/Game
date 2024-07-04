package repository

import (
	"context"
	"fmt"
	"game/app"
	"game/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ShowAllUsers(number int64) ([]domain.InternalUser, error) {
	var userScoreboard []domain.InternalUser

	collection, err := app.Collection()

	if err != nil {
		log.Println("Failed to connect to Db : ", err)
		return userScoreboard, fmt.Errorf("failed to connect to database: %w", err)
	}

	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"point", -1}}).SetLimit(number).SetProjection(bson.D{{"phone", 0}, {"_id", 0}})
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return userScoreboard, err
	}
	err = cursor.All(context.TODO(), &userScoreboard)
	if err != nil {
		return userScoreboard, err
	}
	return userScoreboard, nil
}
