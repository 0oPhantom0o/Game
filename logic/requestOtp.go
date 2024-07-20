package logic

import (
	"fmt"
	"game/constants"
	"game/repository"
)

func RequestOtp(phone string) (string, error) {
	//limiting request otp
	counter, rateLimit := repository.OtpLimit(phone)
	if counter == 1 {
		err := repository.ExpireOtpTime(rateLimit)
		if err != true {
			return "", fmt.Errorf("error in expire time set")
		}
	}
	if counter >= 5 {
		return "", fmt.Errorf("user Limited")
	}

	code, err := randomCode()
	if err != nil {
		return "", err
	}
	err = tempUser(phone, code)
	if err != nil {
		return "", err
	}

	response := constants.CodeIsSent + phone

	return response, nil
}

func tempUser(phone, code string) error {

	err := repository.InsertOtp(phone, code)
	if err != nil {
		return fmt.Errorf("couldnt insert user in temp database")
	}
	return nil
}
