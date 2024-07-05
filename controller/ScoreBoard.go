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
	id := c.Param("number")
	number, err := logic.ConvertStringToInteger(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "number is not valid")
		return
	}
	userList, err := logic.ScoreBoard(number)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	scoreBoard, err := logic.ConvertStructToString(userList)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, scoreBoard)
}
