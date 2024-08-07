package logic

import (
	"fmt"
	"game/constants"
)

func (g *GameLogic) GenerateUser(phone, randomCode string) (string, error) {
	//limiting wronged answers
	status, err := g.checkOtp(phone, randomCode)
	if err != nil {
		return "", err
	}

	if !status {
		counter, rateLimit := g.Repo.OTPAnswerLimit(phone)
		if counter == constants.WrongedAnswerOtpBase {
			err := g.Repo.ExpireWrongedAnswerTime(rateLimit)
			if err != true {
				return "", fmt.Errorf("error in expire time set")
			}
		}
		if counter == constants.WrongedAnswerOtpLimit {
			return "", fmt.Errorf("user Limited. wait 10 minutes")
		}

		return "", fmt.Errorf("wronge answer")
	}
	id, err := g.Repo.CreateUser(phone)
	if err != nil {
		return "", err
	}
	token, err := GenerateToken(id)
	if err != nil {
		return "", fmt.Errorf("failed to generate token")
	}
	return token, nil
}
