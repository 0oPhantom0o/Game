package repository

import (
	"context"
	"fmt"
	"game/app"
	"game/constants"
)

func InsertOtp(key, value string) error {
	var rdb = app.RedisDB
	var ctx = context.Background()

	_, err := rdb.Set(ctx, key, value, constants.TempUserExpireTIme).Result()
	if err != nil {
		return fmt.Errorf("couldnt insert otp in redis database")

	}

	return nil
}
func InsertAnswer(id, answer string) error {
	var rdb = app.RedisDB
	var ctx = context.Background()

	_, err := rdb.Set(ctx, id, answer, constants.QuestionExpireTime).Result()
	if err != nil {
		return fmt.Errorf("couldnt insert question in redis database")
	}

	return nil
}
