package repository

import (
	"context"
	"game/app"
)

func FindUserTempData(id string) (string, error) {
	rdb := app.RedisDB
	ctx := context.Background()
	code, err := rdb.Get(ctx, id).Result()
	if err != nil {
		return "", err
	}
	return code, nil
}
