package logic

import (
	"game/domain"
	"game/repository"
)

func ScoreBoard(number int64) ([]domain.UserScoreBoard, error) {
	data, err := repository.ShowAllUsers(number)
	if err != nil {
		return data, err
	}

	return data, nil
}
