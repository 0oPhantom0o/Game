package repository

import (
	"context"
	"fmt"
	"game/app"
	"game/constants"
	"time"
)

func RedisDataSet(key, value, dataType string) error {
	rdb := app.RedisDB
	ctx := context.Background()
	if dataType == "" {
		_, err := rdb.Set(ctx, key, value, 10*time.Minute).Result()
		if err != nil {
			return fmt.Errorf("couldnt insert data in redis database")

		}
	} else {
		_, err := rdb.Set(ctx, key, value, constants.RedisExpireTime).Result()
		if err != nil {
			return fmt.Errorf("couldnt insert data in redis database")
		}
	}
	return nil
}
