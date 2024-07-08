package repository

func FindStoredOtp(id string) (string, error) {

	value, err := Rdb.Get(Ctx, id).Result()
	if err != nil {
		return "", err
	}
	return value, nil

}
func FindAnswer(id string) (string, error) {

	value, err := Rdb.Get(Ctx, id).Result()
	if err != nil {
		return "", err
	}
	return value, nil

}
