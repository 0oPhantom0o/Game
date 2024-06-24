package logic

import (
	"game/domain"
	"game/repository"
)

func GenerateUser(user *domain.User) (string, error) {
	id, err := repository.CreateUser(user)
	if err != nil {
		return "", err
	}
	return id, nil
}
