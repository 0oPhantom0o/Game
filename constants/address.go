package constants

const (
	ServerPort = "localhost:8080"
	RedisPort  = "localhost:6380"
	MongoPort  = "mongodb://localhost:27017"
)

const (
	Database       = "Game"
	UserCollection = "User"
)
const DatabaseInitError = "Failed to initialize MongoDB: %v"
