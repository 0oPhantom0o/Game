package app

import (
	"fmt"
)

func InitDb() error {
	err := mongoConnection()
	if err != nil {
		return fmt.Errorf("failed to initialize MongoDB:%v", err)
	}
	err = redisConnection()
	if err != nil {
		return fmt.Errorf("failed to initialize RedisDB:%v", err)
	}

	return nil

}
