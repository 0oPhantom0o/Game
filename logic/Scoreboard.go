package logic

import (
	"game/repository"
	"go.mongodb.org/mongo-driver/bson"
)

func ScoreBoard(number int64) ([]bson.D, error) {
	data, err := repository.ShowAllUsers(number)
	if err != nil {
		return nil, err
	}

	return data, nil
}
