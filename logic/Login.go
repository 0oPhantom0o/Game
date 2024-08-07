package logic

func (g *GameLogic) Login(phone string) (string, error) {
	id, err := g.Repo.FindUserIdByPhone(phone)
	if err != nil {
		return "", err
	}
	token, err := GenerateToken(id)
	if err != nil {

		return "", err
	}
	return token, nil

}
