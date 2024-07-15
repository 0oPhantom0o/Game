package controller

import (
	"game/domain"
	"game/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

// a
func Answer(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	id, err := logic.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token has problem"})
		return
	}
	var user domain.UserAnswer
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status, err := logic.CheckAnswer(id, user.Answer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if status == true {
		c.JSON(http.StatusOK, gin.H{"message": "right answer"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "wronged answer"})
}
