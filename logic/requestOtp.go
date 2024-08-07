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
			return fmt.Errorf("error in expire time set")
		}
	}
	if counter >= 5 {
		return fmt.Errorf("user Limited")
	}

	code, err := randomCode()
	if err != nil {
		return err
	}
	err = g.tempUser(phone, code)
	if err != nil {
		return err
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
