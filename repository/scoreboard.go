package repository

import (
	"game/constants"
	"game/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *ConRepository) ShowUsers(page, limit int64) ([]domain.TopPlayers, error) {
	var scoreboard []domain.TopPlayers
	collection := repo.mongodb.Database(constants.Database).Collection(constants.UserCollection)

	filter := bson.D{}
	skip := (page - 1) * limit

	opts := options.Find().SetSort(bson.D{{"point", -1}}).
		SetLimit(limit + constants.DatabaseNextPageCheck).SetSkip(skip)
	cursor, err := collection.Find(repo.ctx, filter, opts)
	if err != nil {
		return scoreboard, err
	}
	err = cursor.All(repo.ctx, &scoreboard)
	if err != nil {
		return scoreboard, err
	}

	return scoreboard, nil
}
