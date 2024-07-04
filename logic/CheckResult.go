package logic

import (
	"fmt"
	"game/repository"
)

func CheckAnswer(id, answer string) (bool, error) {
	result, err := repository.FindRedisValue(id)
	if err != nil {
		return false, fmt.Errorf("too late... find new question ")
	}
	if result == answer {
		err = ChangePoint(id, 1)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	if result != answer {
		err = ChangePoint(id, -1)
		if err != nil {
			return false, err
		}
		return false, nil
	}
	return false, fmt.Errorf("answer is wronge")
}
