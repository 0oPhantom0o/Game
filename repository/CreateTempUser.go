package repository

import (
	"context"
	"fmt"
	"game/app"
	"game/constants"
	"time"
)

func TempUser(phone, code string) error {
	rdb := app.RedisDB
	ctx := context.Background()
	_, err := rdb.Set(ctx, phone, code, constants.RedisExpireTime).Result()
	if err != nil {
		return fmt.Errorf("couldnt insert data in redis database")
	}
	return nil
}
func UserResult(id, answer string) error {
	rdb := app.RedisDB
	ctx := context.Background()
	_, err := rdb.Set(ctx, id, answer, 10*time.Minute).Result()
	if err != nil {
		return err
	}
	return nil
}
