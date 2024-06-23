package logic

import "game/repository"

func NumberForUser(id string) (string, error) {
	result := "4"
	err := repository.UserResult(id, result)
	if err != nil {
		return "", nil
	}
	return "2+2: ", nil
}
