package logic

import (
	"game/constants"
	"game/domain"
	"game/repository"
)

func ScoreBoard(number, limit string) ([]domain.TopPlayers, bool, error) {

	page, err := convertStringToInteger(number)
	count, err := convertStringToInteger(limit)

	if err != nil {
		return nil, false, err
	}
	scoreBoard, err := repository.ShowUsers(page, count)
	scoreBoard, nextPage := checkNextPage(scoreBoard, count)
	if err != nil {
		return nil, nextPage, err
	}

	return scoreBoard, nextPage, nil

}

func checkNextPage(scoreboard []domain.TopPlayers, count int64) ([]domain.TopPlayers, bool) {

	if len(scoreboard) > int(count) {
		scoreboard = scoreboard[:len(scoreboard)-constants.RemoveNextPageCheck]
		return scoreboard, true
	}
	//a
	return scoreboard, false
}
