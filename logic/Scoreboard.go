package logic

import (
	"game/repository"
	"go.mongodb.org/mongo-driver/bson"
)

func ScoreBoard() ([]bson.D, error) {
	data, err := repository.ShowAllUsers()
	if err != nil {
		return nil, err
	}

	return data, nil
}
