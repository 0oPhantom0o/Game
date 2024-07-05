package constants

import "time"

const (
	QuestionExpireTime      = 30 * time.Second
	TempUserExpireTIme      = 5 * time.Minute
	RateLimitUserExpireTime = 10 * time.Minute
)
