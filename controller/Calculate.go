package controller

import (
	"game/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Calculate(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, "error")
		return
	}
	Id, err := logic.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "error")
		return
	}
	number, err := logic.NumberForUser(Id)
	c.JSON(http.StatusOK, number)

}
