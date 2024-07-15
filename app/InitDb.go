package app

import (
	"fmt"
)

// a
func InitDb() error {
	err := MongoConnection()
	if err != nil {
		return fmt.Errorf("failed to initialize MongoDB: %v", err)
	}
	err = RedisConnection()
	if err != nil {
		return fmt.Errorf("failed to initialize RedisDB: %v", err)
	}

	return nil

}
