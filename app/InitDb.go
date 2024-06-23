package app

import (
	"game/constants"
	"log"
)

func InitDb() error {
	err := MongoConnection()
	if err != nil {
		log.Fatalf(constants.DatabaseInitError, err)
	}
	err = RedisConnection()
	if err != nil {
		log.Fatalf(constants.DatabaseInitError, err)
	}

	return nil

}
