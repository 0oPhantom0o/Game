package logic

import (
	"fmt"
	"game/constants"
)

func (g *GameLogic) CheckAnswer(id, answer string) (bool, error) {
	// Use g.redisDB instead of repository directly
	mongoId, err := convertStringToPrimitive(id)
	if err != nil {
		return false, err
	}
	trueAnswer, err := g.Repo.FindStoredAnswer(id)
	if err != nil {
		return false, fmt.Errorf("wrong answer. search for a new question")
	}

	if trueAnswer == answer {
		err := g.Repo.ChangePoint(mongoId, constants.RightAnswerPoint)
		if err != nil {
			return true, err
		}
		return true, nil
	} else {
		err = g.Repo.ChangePoint(mongoId, constants.WrongedAnswerPoint)
		if err != nil {
			return false, err
		}
	}
	err = g.Repo.DeleteAnswer(id)
	if err != nil {
		return true, err
	}
	return false, nil
}

// Similarly, update other logic functions to use the interfaces.

func (g *GameLogic) changePoint(id string, point int) error {
	mongoId, err := convertStringToPrimitive(id)
	if err != nil {
		return err
	}
	err = g.Repo.ChangePoint(mongoId, point)
	if err != nil {
		return err
	}
	return nil
}
