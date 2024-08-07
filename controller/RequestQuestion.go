package controller

import (
	"game/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctrl *GameController) RequestQuestion(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")
	Id, err := logic.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	question, err := ctrl.Logic.Question(Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"answer_question:": question})

}
