package logic

import "game/repository"

type GameLogic struct {
	Repo repository.Repository
}

func NewRestaurantService(repo repository.Repository) *GameLogic {
	return &GameLogic{Repo: repo}
}
