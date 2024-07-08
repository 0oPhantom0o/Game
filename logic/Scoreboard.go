package logic

import (
	"game/domain"
	"game/repository"
)

func ScoreBoard(number, limit string) ([]domain.TopPlayers, error) {

	page, err := convertStringToInteger(number)
	count, err := convertStringToInteger(limit)

	if err != nil {
		return nil, err
	}
	scoreBoard, err := repository.ShowUsers(page, count)
	if err != nil {
		return nil, err
	}

	return scoreBoard, nil

}
