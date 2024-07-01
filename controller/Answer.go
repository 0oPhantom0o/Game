package controller

import (
	"game/domain"
	"game/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Result(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	id, err := logic.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "error")
		return
	}
	var Answer domain.Answer
	if err := c.BindJSON(&Answer); err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}

	status, err := logic.CheckAnswer(id, Answer.Result)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}

	c.JSON(http.StatusOK, status)
}
