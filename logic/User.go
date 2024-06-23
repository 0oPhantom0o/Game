package logic

import (
	"game/domain"
	"game/repository"
)

func GenerateUser(user *domain.User) (string, error) {
	Id, err := repository.CreateUser(user)
	if err != nil {
		return "", err
	}
	return Id, nil
}
