package repository

import (
	"fmt"
)

func (repo *ConRepository) DeleteAnswer(id string) error {

	_, err := repo.redisdb.Del(repo.ctx, id).Result()
	if err != nil {
		return fmt.Errorf("failed to delete OTP from Redis: %w", err)
	}
	return nil
}
