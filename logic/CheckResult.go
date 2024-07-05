package logic

import (
	"fmt"
	"game/repository"
)

func CheckAnswer(id, answer string) (bool, error) {
	result, err := repository.FindRedisValue(id)
	if err != nil {
		return false, fmt.Errorf("too late... please search new question ")
	}
	if result == answer {
		return true, nil
	}
	return false, nil
}
