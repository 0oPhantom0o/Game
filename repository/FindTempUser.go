package repository

import (
	"context"
	"game/app"
)

func FindRandCode(user string) (string, error) {
	rdb := app.RedisDB
	ctx := context.Background()
	code, err := rdb.Get(ctx, user).Result()
	if err != nil {
		return "", err
	}
	return code, nil
}
func FindUserResult(id string) (string, error) {
	rdb := app.RedisDB
	ctx := context.Background()
	code, err := rdb.Get(ctx, id).Result()
	if err != nil {
		return "", err
	}
	return code, nil
}
