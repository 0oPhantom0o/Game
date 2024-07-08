package repository

import (
	"context"
	"fmt"
	"game/app"
	"game/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ShowUsers(page, limit int64) ([]domain.TopPlayers, error) {
	var userScoreboard []domain.TopPlayers
	collection, err := app.Collection()
	if err != nil {
		return userScoreboard, fmt.Errorf("failed to connect to database: %w", err)
	}
	filter := bson.D{}
	//	limit := int64(10)
	skip := (page - 1) * limit
	//paginatedData, err := New(collection).Context(ctx).Limit(limit).Page(page).Sort("price", -1).Select(projection).Filter(filter).Decode(&products).Find()
	//only get nickname and point
	//SetProjection(bson.D{{"phone", 0}, {"_id", 0}, {"nickNameLimit", 0}})
	opts := options.Find().SetSort(bson.D{{"point", -1}}).
		SetLimit(limit).SetSkip(skip)
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
