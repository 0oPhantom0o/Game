package controller

import (
	"game/domain"
	"game/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request error")
		return
	}
	id, err := logic.FindUser(user.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, "failed to find user")
		return
	}

	token, err := logic.GenerateToken(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "failed to generate token")

		return
	}
	c.JSON(http.StatusOK, token)
}
