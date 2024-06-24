package logic

import "game/repository"

func ChangePoint(id string, point int) error {
	mongoId, err := ConvertStringToPrimivite(id)
	if err != nil {
		return err
	}
	userData, err := repository.FindUser(mongoId)
	if err != nil {
		return err
	}
	userData.Point = point
	err = repository.ChangePoint(userData, mongoId)
	if err != nil {
		return err
	}
	return nil
}
