package repository

import (
	"context"
	"game/app"
	"time"
)

func TempUser(phone, code string) error {
	rdb := app.RedisDB
	ctx := context.Background()
	_, err := rdb.Set(ctx, phone, code, 10*time.Minute).Result()
	if err != nil {
		return err
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
