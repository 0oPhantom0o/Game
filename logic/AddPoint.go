package logic

import "game/repository"

func AddPoint(id string) error {
	mongoId, err := ConvertStringToPrimivite(id)
	if err != nil {
		return err
	}
	use, err := repository.FindUser(mongoId)
	if err != nil {
		return err
	}

	err = repository.IncreasePoint(use)
	if err != nil {
		return err
	}
	return nil
}
