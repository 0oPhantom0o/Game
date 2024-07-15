package logic

import "game/repository"

func Login(phone string) (string, error) {
	id, err := repository.FindUserIdByPhone(phone)
	if err != nil {
		return "", err
	}
	token, err := GenerateToken(id)
	if err != nil {

		return "", err
	}
	return token, nil

}

//a
