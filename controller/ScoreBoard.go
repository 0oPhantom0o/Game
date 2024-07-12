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
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}
	number := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("count", "10")
	scoreBoard, nextPage, err := logic.ScoreBoard(number, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !nextPage {
		c.JSON(http.StatusOK, gin.H{"players": scoreBoard, "next page": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"players": scoreBoard, "next page": true})
}
