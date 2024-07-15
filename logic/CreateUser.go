package logic

import (
	"fmt"
	"game/constants"
	"game/repository"
)

func GenerateUser(phone, randomCode string) (string, error) {
	//limiting wronged answers
	status, err := checkOtp(phone, randomCode)
	if err != nil {
		return "", err
	}

	if !status {
		limitCounter, PhoneLimit := repository.OTPAnswerLimit(phone)
		if limitCounter == constants.WrongedAnswerOtpBase {
			err := repository.ExpireWrongedAnswerTime(PhoneLimit)
			if err != true {
				return "", fmt.Errorf("error in expire time set")
			}
		}
		if limitCounter == constants.WrongedAnswerOtpLimit {
			return "", fmt.Errorf("user Limited. wait 10 minutes")
		}

		return "", fmt.Errorf("wronge answer")
	}
	//a
	id, err := repository.CreateUser(phone)
	if err != nil {
		return "", err
	}
	token, err := GenerateToken(id)
	if err != nil {
		return "", fmt.Errorf("failed to generate token")
	}
	return token, nil
}
