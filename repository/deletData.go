package repository

import (
	"context"
	"fmt"
	"game/app"
)

func DeleteAnswer(id string) error {
	rdb := app.RedisDB
	ctx := context.Background()

	_, err := rdb.Del(ctx, id).Result()
	if err != nil {
		return fmt.Errorf("failed to delete OTP from Redis: %w", err)
	}
	return nil
}
