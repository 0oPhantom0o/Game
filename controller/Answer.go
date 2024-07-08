package controller

import (
	"game/domain"
	"game/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Answer(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	id, err := logic.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "error")
		return
	}
	var user domain.UserAnswer
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}

	status, err := logic.CheckAnswer(id, user.Answer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}
	if status == true {
		c.JSON(http.StatusOK, "right answer")
		return
	}
	c.JSON(http.StatusOK, "wronged answer")
}
