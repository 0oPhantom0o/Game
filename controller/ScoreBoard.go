package controller

import (
	"game/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctrl *GameController) ScoreBoard(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	_, err := logic.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	pageNumber := c.DefaultQuery("page", "1")
	perPage := c.DefaultQuery("count", "10")
	score, err := ctrl.Logic.ScoreBoard(pageNumber, perPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, score)
}
