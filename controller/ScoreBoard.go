package controller

import (
	"game/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ScoreBoard(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	_, err := logic.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "error")
		return
	}
	listOfUsers, err := logic.ScoreBoard()
	scoreBoard, err := logic.ConvertBsonDToScoreBoard(listOfUsers)
	c.JSON(http.StatusOK, scoreBoard)
}
