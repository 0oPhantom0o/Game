package logic

import (
	"game/domain"
	"game/repository"
)

func ScoreBoard(number string) ([]domain.InternalUser, error) {
	count, err := convertStringToInteger(number)
	if err != nil {
		return nil, err
	}
	scoreBoard, err := repository.ShowAllUsers(count)
	if err != nil {
		return nil, err
	}

	return scoreBoard, nil

}
