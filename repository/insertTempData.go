package repository

import (
	"fmt"
	"game/constants"
)

func (repo *ConRepository) InsertOtp(phone, code string) error {

	_, err := repo.redisdb.Set(repo.ctx, phone, code, constants.TempUserExpireTIme).Result()
	if err != nil {
		return fmt.Errorf("couldnt insert otp in redis database")

	}

	return nil
}
func (repo *ConRepository) InsertAnswer(id, answer string) error {

	_, err := repo.redisdb.Set(repo.ctx, id, answer, constants.QuestionExpireTime).Result()
	if err != nil {
		return fmt.Errorf("couldnt insert question in redis database")
	}

	return nil
}
