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
	var User domain.User
	if err := c.BindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}

	status, err := logic.CheckAnswer(id, User.Result)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}

	c.JSON(http.StatusOK, status)
}
