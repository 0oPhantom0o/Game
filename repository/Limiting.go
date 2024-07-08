package repository

func OTPLimit(phone string) (int64, string) {

	//insert rate limit phone and add counter
	rateLimit := "rateLimit : " + phone
	count := Rdb.Incr(Ctx, rateLimit).Val()

	return count, rateLimit
}
func OTPAnswerLimit(phone string) (int64, string) {

	//insert rate limit phone and add counter
	rateLimit := "WrongedAnswerLimit : " + phone
	count := Rdb.Incr(Ctx, rateLimit).Val()

	return count, rateLimit
}
