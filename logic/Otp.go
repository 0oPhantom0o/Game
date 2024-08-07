package logic

func (g *GameLogic) checkOtp(phone, randomCode string) (bool, error) {
	storedCode, err := g.Repo.FindStoredOtp(phone)
	if err != nil {
		return false, err
	}
	if storedCode == randomCode && storedCode != "" {
		return true, nil
	}

	return false, nil

}
