package controller

import (
	"game/domain"
	"game/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctrl *GameController) Answer(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	id, err := logic.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	var user domain.RequestAnswer
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	status, err := ctrl.Logic.CheckAnswer(id, user.Answer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if !status {
		c.JSON(http.StatusOK, gin.H{"message": "wronged_answer"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "right_answer"})
}
