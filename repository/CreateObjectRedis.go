package repository

import (
	"context"
	"fmt"
	"game/app"
	"game/constants"
)

func RedisDataSet(key, value, dataType string) error {
	rdb := app.RedisDB
	ctx := context.Background()

	if dataType == "TempUser" {
		_, err := rdb.Set(ctx, key, value, constants.TempUserExpireTIme).Result()
		if err != nil {
			return fmt.Errorf("couldnt insert TempUser in redis database")

		}
	} else if dataType == "question" {
		_, err := rdb.Set(ctx, key, value, constants.QuestionExpireTime).Result()
		if err != nil {
			return fmt.Errorf("couldnt insert question in redis database")
		}
	}
	return nil
}
