package logic

import (
	"fmt"
	"game/constants"
	"game/repository"
)

func CheckAnswer(id, answer string) (bool, error) {

	trueAnswer, err := repository.FindAnswer(id)
	if err != nil {
		return false, fmt.Errorf("wronge answer . search for new question ")
	}

	if trueAnswer == answer {

		err := changePoint(id, constants.RightAnswerPoint)
		if err != nil {
			return true, err
		}
		return true, nil
	} else {
		err = changePoint(id, constants.WrongedAnswerPoint)
		if err != nil {
			return false, err
		}
	}
	err = repository.DeleteAnswer(id)
	if err != nil {
		return true, err
	}
	return false, nil
}

func changePoint(id string, point int) error {
	mongoId, err := convertStringToPrimitive(id)
	if err != nil {
		return err
	}
	err = repository.ChangePoint(mongoId, point)
	if err != nil {
		return err
	}
	return nil
}

//a
