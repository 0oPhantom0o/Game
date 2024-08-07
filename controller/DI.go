package controller

import "game/logic"

type GameController struct {
	Logic *logic.GameLogic
}

func NewGameController(service *logic.GameLogic) *GameController {
	return &GameController{Logic: service}
}
