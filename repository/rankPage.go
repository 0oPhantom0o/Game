package repository

import (
	"context"
	"fmt"
	"game/app"
	"game/constants"
	"game/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ShowUsers(page, limit int64) ([]domain.TopPlayers, error) {
	var scoreboard []domain.TopPlayers
	collection, err := app.Collection()
	if err != nil {
		return scoreboard, fmt.Errorf("failed to connect to database: %w", err)
	}
	filter := bson.D{}
	//	limit := int64(10)
	skip := (page - 1) * limit
	//paginatedData, err := New(collection).Context(ctx).Limit(limit).Page(page).Sort("price", -1).Select(projection).Filter(filter).Decode(&products).Find()
	//only get nickname and point
	//SetProjection(bson.D{{"phone", 0}, {"_id", 0}, {"nickNameLimit", 0}})
	opts := options.Find().SetSort(bson.D{{"point", -1}}).
		SetLimit(limit + constants.DatabaseNextPageCheck).SetSkip(skip)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return scoreboard, err
	}
	err = cursor.All(context.TODO(), &scoreboard)
	if err != nil {
		return scoreboard, err
	}

	return scoreboard, nil
}
