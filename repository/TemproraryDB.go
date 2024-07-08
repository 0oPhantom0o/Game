package repository

import (
	"context"
	"fmt"
	"game/app"
	"game/constants"
)

var Rdb = app.RedisDB
var Ctx = context.Background()

func InsertOtp(key, value string) error {
	_, err := Rdb.Set(Ctx, key, value, constants.TempUserExpireTIme).Result()
	if err != nil {
		return fmt.Errorf("couldnt insert otp in redis database")

	}

	return nil
}
func InsertAnswer(id, answer string) error {

	_, err := Rdb.Set(Ctx, id, answer, constants.QuestionExpireTime).Result()
	if err != nil {
		return fmt.Errorf("couldnt insert question in redis database")
	}

	return nil
}
