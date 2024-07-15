package logic

import (
	"fmt"
	"game/repository"
)

func RequestOtp(phone string) error {
	//limiting request otp
	limitCounter, rateLimit := repository.OTPLimit(phone)
	if limitCounter == 1 {
		err := repository.ExpireOtpTime(rateLimit)
		if err != true {
			return fmt.Errorf("error in expire time set")
		}
	}
	if limitCounter >= 5 {
		return fmt.Errorf("user Limited")
	}

	code, err := RandomCode()
	if err != nil {
		return err
	}
	err = tempUser(phone, code)
	if err != nil {
		return err
	}
	return nil
}

// a
func tempUser(phone, code string) error {

	err := repository.InsertOtp(phone, code)
	if err != nil {
		return fmt.Errorf("couldnt insert user in temp database")
	}
	return nil
}
