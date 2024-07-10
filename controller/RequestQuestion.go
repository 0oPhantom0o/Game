package controller

import (
	"game/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequestQuestion(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	Id, err := logic.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "error")
		return
	}
	question, err := logic.Question(Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}
	c.JSON(http.StatusOK, question)

}
