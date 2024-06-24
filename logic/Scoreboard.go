package logic

import (
	"game/repository"
)

func ScoreBoard() error {
	err := repository.ShowAllUsers2()
	if err != nil {
		return err
	}
	return nil
}
