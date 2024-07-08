package controller

import (
	"game/domain"
	"game/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOtp(c *gin.Context) {

	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, "bad request error")
		return
	}

	token, err := logic.GenerateUser(user.Phone, user.RandomCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, "failed to create user")
		return
	}

	c.JSON(http.StatusOK, token)

}
