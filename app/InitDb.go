package app

import (
	"log"
)

func InitDb() error {
	err := MongoConnection()
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB: %v", err)
	}
	err = RedisConnection()
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB: %v", err)
	}

	return nil

}
