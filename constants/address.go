package constants

const (
	ServerPort = "localhost:8080"
	RedisPort  = "localhost:6380"
	MongoPort  = "mongodb://localhost:27017"
)

const (
	// Database mongo collection name
	Database       = "Game"
	UserCollection = "User"
)

// router groups
const (
	Version = "/v1"
	//Auth routs
	Auth       = "/auth"
	RequestOtp = "/sign_up"
	Submit     = "/submitOtp"
	Login      = "/login"
	Nickname   = "/nickname"

	//Game routs
	Game            = "/game"
	RequestQuestion = "/question"
	SubmitAnswer    = "/answer"
	ShowTopPlayers  = "/scoreboard/"
)
