package logic

import (
	"fmt"
	"game/repository"
)

func CheckAnswer(id, answer string) error {
	result, err := repository.FindUserResult(id)
	if err != nil {
		return err
	}
	if result == answer {

		return nil
	}
	return fmt.Errorf("answer is wronge")
}
