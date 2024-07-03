package repository

import (
	"context"
	"game/app"
)

func FindUserValue(id string) (string, error) {
	rdb := app.RedisDB
	ctx := context.Background()
	value, err := rdb.Get(ctx, id).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}
