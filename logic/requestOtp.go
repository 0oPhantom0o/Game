package logic

import (
	"fmt"
)

func (g *GameLogic) RequestOtp(phone string) error {
	//limiting request otp
	counter, rateLimit := g.Repo.OtpLimit(phone)
	if counter == 1 {
		err := g.Repo.ExpireOtpTime(rateLimit)
		if err != true {
			return fmt.Errorf("failed to set expiration time: %v", err)
		}
	}
	if counter >= 5 {
		return fmt.Errorf("user has reached the maximum limit of OTP requests")
	}

	code, err := randomCode()
	if err != nil {
		return fmt.Errorf("failed to generate random code: %v", err)
	}
	err = g.tempUser(phone, code)
	if err != nil {
		return fmt.Errorf("failed to store temporary user: %v", err)
	}

	return nil
}

func (g *GameLogic) tempUser(phone, code string) error {

	err := g.Repo.InsertOtp(phone, code)
	if err != nil {
		return fmt.Errorf("couldnt insert user in temp database")
	}
	return nil
}
