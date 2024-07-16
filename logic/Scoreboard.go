package logic

import (
	"game/constants"
	"game/domain"
	"game/repository"
)

func ScoreBoard(number, limit string) (domain.Score, error) {
	var score domain.Score

	page, err := convertStringToInteger(number)
	if err != nil {
		return score, err
	}
	count, err := convertStringToInteger(limit)

	if err != nil {
		return score, err
	}
	scoreBoard, err := repository.ShowUsers(page, count)
	score.Players, score.NextPage = checkNextPage(scoreBoard, count)
	if err != nil {
		return score, err
	}

	return score, nil

}

func checkNextPage(scoreboard []domain.TopPlayers, count int64) ([]domain.TopPlayers, bool) {

	if len(scoreboard) > int(count) {

		scoreboard = scoreboard[:len(scoreboard)-constants.DatabaseNextPageCheck]
		return scoreboard, true
	}

	return scoreboard, false
}
